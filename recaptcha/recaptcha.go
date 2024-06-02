package recaptcha

import (
	"encoding/json"
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

func (service RecaptchaService) Verify(token string) bool {
	data := url.Values{}
	data.Set("secret", service.secret)
	data.Set("response", token)

	resp, err := http.PostForm(service.verifyLink, data)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false
	}
	if success, ok := result["success"].(bool); !ok || !success {
		return false
	}
	if score, ok := result["score"].(float64); !ok || score < service.boundryRating {
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
