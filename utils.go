package main

import (
	"bytes"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/types/api"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

func validatePayloadRequestPayload(payload api.RequestAppointmentPayload) bool {
	if payload.Email == "" {
		return false
	}
	if payload.Name == "" {
		return false
	}
	if payload.PhoneNumber == "" {
		return false
	}
	isValid := appServices.RecaptchaService.Verify(payload.RecaptchaToken)
	return isValid
}

func GetHTMLFromRootTemplate(name string) *bytes.Buffer {
	byt, err := web.GetHTMLFromTemplate(html, name)
	if err != nil {
		return new(bytes.Buffer)
	}

	return &byt
}
