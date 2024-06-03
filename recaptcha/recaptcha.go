package recaptcha

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type RecaptchaService struct {
	verifyLink    string
	secret        string
	boundryRating float64
}

type RecaptchaResponse struct {
	Success bool    `json:"success"`
	Score   float64 `json:"score"`
}

func (service RecaptchaService) Verify(token string) bool {
	data := url.Values{}
	data.Set("secret", service.secret)
	data.Set("response", token)

	resp, err := http.PostForm(service.verifyLink, data)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	bodyByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var recaptchaResponse RecaptchaResponse
	err = json.Unmarshal(bodyByte, &recaptchaResponse)
	if err != nil {
		return false
	}
	if !recaptchaResponse.Success || recaptchaResponse.Score < service.boundryRating {
		return false
	}

	return true
}

func CreateRecaptchaService() *RecaptchaService {
	key := os.Getenv("RECAPTCHA_KEY")
	if key == "" {
		log.Fatal("Needs recaptcha key")
	}
	url := os.Getenv("RECAPTCHA_VERIFY_LINK")
	if url == "" {
		log.Fatal("Needs link of where to verify recaptcha link")
	}
	return &RecaptchaService{
		secret:        key,
		verifyLink:    url,
		boundryRating: .6,
	}
}
