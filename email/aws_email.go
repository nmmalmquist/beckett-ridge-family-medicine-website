package email

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/types/api"
)

type EmailService struct {
	ses                *sesv2.Client
	fromEmail          string
	requestApptToEmail string
}

func (service EmailService) send(subject string, body string, to string) (*sesv2.SendEmailOutput, error) {
	input := &sesv2.SendEmailInput{
		FromEmailAddress: aws.String(service.fromEmail),
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: aws.String(subject),
				},
				Body: &types.Body{
					Text: &types.Content{
						Data: aws.String(body),
					},
				},
			},
		},
	}
	return service.ses.SendEmail(context.TODO(), input)
}

func (service EmailService) SendAppointmentRequest(payload api.RequestAppointmentPayload) (*sesv2.SendEmailOutput, error) {
	return service.send(fmt.Sprintf("%s requests an appointment", payload.Name),
		fmt.Sprintf("Please call %s to create an appointment with them\nReason: %s",
			payload.PhoneNumber, payload.Text), service.requestApptToEmail)
}

func CreateEmailService() *EmailService {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	from := os.Getenv("FROM_EMAIL")
	if from == "" {
		log.Fatal("The from address for emails is not specified")
	}
	to := os.Getenv("REQUEST_APPT_TO_EMAIL")
	if to == "" {
		log.Fatal("The to address for request appointment emails is not specified")
	}
	return &EmailService{
		ses:                sesv2.NewFromConfig(cfg),
		fromEmail:          from,
		requestApptToEmail: to,
	}
}
