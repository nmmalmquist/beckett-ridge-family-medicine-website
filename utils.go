package main

import (
	"nickmalmquist.com/beckett-ridge-family-medicine-website/types/api"
)

func validatePayloadRequestPayload(payload api.RequestAppointmentPayload) bool{
if (payload.Email == ""){
	return false
}
if (payload.Name == ""){
	return false
}
if (payload.PhoneNumber == ""){
	return false
}
isValid := recaptchaService.Verify(payload.RecaptchaToken)
return isValid

}
