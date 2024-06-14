package main

import (
	"log"
	"net/http"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/constants"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/types"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/types/api"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

func index(r *http.Request) *web.Response {
	// Need to check because all routes that don't match will come here
	if r.URL.Path != "/" {
		return error404()
	}
	type IndexPageData struct {
		ActivePage string
		CardFacts  []types.CardFact
		TabFacts   []types.TabFact
		Articles   []types.Article
	}

	data := IndexPageData{
		ActivePage: "index",
		CardFacts: []types.CardFact{
			{Icon: GetHTMLFromRootTemplate("components/dr-icon.html"), Title: "Understand Your Health", Description: "At our family practice, we understand that symptoms often hint at deeper underlying issues. We prioritize thorough examination and personalized care to uncover these root causes, ensuring a comprehensive understanding of your health. By taking the time to explore these factors, we craft personalized health plans tailored to your needs, providing you with a clear picture of your health journey. Invest in comprehensive care today for a healthier tomorrow."},
			{Icon: GetHTMLFromRootTemplate("components/currency-circle-icon.html"), Title: "Save Money", Description: "By exploring the underlying health issues with a trusted family practice, you not only invest in your well-being but also save significant amounts of money in the long run. Our comprehensive approach to healthcare focuses on preventative measures, early detection, and personalized treatment plans tailored to your unique needs. By addressing health concerns proactively, we help you avoid costly medical emergencies and unnecessary expenses associated with untreated conditions. Invest in your health today to save money and enjoy a healthier, more fulfilling life tomorrow."},
			{Icon: GetHTMLFromRootTemplate("components/vitality-icon.html"), Title: "Live With Vitality", Description: "At our family practice, we believe that true health goes beyond the absence of illness; it's about embracing vitality in every aspect of life. We prioritize holistic care that nurtures your physical, mental, and emotional well-being, empowering you to live each day to the fullest. By focusing on preventive measures, healthy lifestyle choices, and personalized wellness plans, we help you unlock your body's full potential and thrive with vitality. Invest in your vitality today for a brighter, more vibrant tomorrow."},
		},
		TabFacts: []types.TabFact{
			{Icon: GetHTMLFromRootTemplate("components/heart-beep-icon.html"), Title: "Healthcare provider for 25+ years", ColorClass: "primary"},
			{Icon: GetHTMLFromRootTemplate("components/award-icon.html"), Title: "180,000+ successful patient visits", ColorClass: "secondary"},
			{Icon: GetHTMLFromRootTemplate("components/people-icon.html"), Title: "More than 2000 happy patients", ColorClass: "primary"},
		},
		Articles: *constants.GetArticlesFromLibrary(),
	}
	return web.HTML(http.StatusOK, html, "pages/index.html", data, nil)
}
func providers(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "pages/providers.html", nil, nil)
}
func privacyPolicy(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "pages/privacy-policy.html", nil, nil)
}
func requestAppointment(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "pages/request-appointment.html", nil, nil)
}
func requestAppointmentPOST(r *http.Request) *web.Response {
	if r.Method != "POST" {
		return web.Data(http.StatusMethodNotAllowed, nil, nil)
	}
	if parseErr := r.ParseForm(); parseErr != nil {
		return web.ErrorJSON(http.StatusBadRequest, "Could not parse form.", nil)
	}

	payload := api.RequestAppointmentPayload{
		Name:           r.FormValue("full-name"),
		Email:          r.FormValue("email"),
		PhoneNumber:    r.FormValue("phone"),
		Text:           r.FormValue("text"),
		RecaptchaToken: r.FormValue("g-recaptcha-response"),
	}

	isValid := validatePayloadRequestPayload(payload)
	if !isValid {
		log.Println("error: payload not valid")
		return GetErrorModal(http.StatusBadRequest)
	}

	_, err := appServices.EmailService.SendAppointmentRequest(payload)

	if err != nil {
		log.Println("error: could not send appointment email -- ", err)
		return GetErrorModal(http.StatusInternalServerError)
	} else {
		return GetSuccessModal()
	}

}
func error404() *web.Response {
	return web.HTML(http.StatusOK, html, "pages/404.html", nil, nil)
}

func robotsTxt(r *http.Request) *web.Response {
	return web.Data(http.StatusOK, []byte(robotsContent), nil)
}
func errorPage(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "pages/error.html", nil, nil)
}
