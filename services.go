package main

import (
	"nickmalmquist.com/beckett-ridge-family-medicine-website/email"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/recaptcha"
)

type AppServices struct {
	// Services
	EmailService     *email.EmailService
	RecaptchaService *recaptcha.RecaptchaService
}

func InitServices() *AppServices {
	return &AppServices{
		EmailService:     email.CreateEmailService(),
		RecaptchaService: recaptcha.CreateRecaptchaService(),
	}
}
