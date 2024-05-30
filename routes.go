package main

import (
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
			{Icon: staticHTML["dr-icon"], Title: "Understand Your Health", Description: "At our family practice, we understand that symptoms often hint at deeper underlying issues. We prioritize thorough examination and personalized care to uncover these root causes, ensuring a comprehensive understanding of your health. By taking the time to explore these factors, we craft personalized health plans tailored to your needs, providing you with a clear picture of your health journey. Invest in comprehensive care today for a healthier tomorrow."},
			{Icon: staticHTML["currency-circle-icon"], Title: "Save Money", Description: "By exploring the underlying health issues with a trusted family practice, you not only invest in your well-being but also save significant amounts of money in the long run. Our comprehensive approach to healthcare focuses on preventative measures, early detection, and personalized treatment plans tailored to your unique needs. By addressing health concerns proactively, we help you avoid costly medical emergencies and unnecessary expenses associated with untreated conditions. Invest in your health today to save money and enjoy a healthier, more fulfilling life tomorrow."},
			{Icon: staticHTML["vitality-icon"], Title: "Live With Vitality", Description: "At our family practice, we believe that true health goes beyond the absence of illness; it's about embracing vitality in every aspect of life. We prioritize holistic care that nurtures your physical, mental, and emotional well-being, empowering you to live each day to the fullest. By focusing on preventive measures, healthy lifestyle choices, and personalized wellness plans, we help you unlock your body's full potential and thrive with vitality. Invest in your vitality today for a brighter, more vibrant tomorrow."},
		},
		TabFacts: []types.TabFact{
			{Icon: staticHTML["heart-beep-icon"], Title: "Healthcare provider for 25+ years", ColorClass: "primary"},
			{Icon: staticHTML["award-icon"], Title: "180,000+ successful patient visits", ColorClass: "secondary"},
			{Icon: staticHTML["people-icon"], Title: "More than 2000 happy patients", ColorClass: "primary"},
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
		Name:        r.FormValue("full-name"),
		Email:       r.FormValue("email"),
		PhoneNumber: r.FormValue("phone"),
		Text:        r.FormValue("text"),
	}
	_, err := emailService.SendAppointmentRequest(payload)
	
	var modalProps types.Modal
	if err != nil {
		modalProps = types.Modal{
			Type:             "error",
			ModalTitle:       "An error occured",
			ModalSubtitle:    "Please try again later",
			ModalIcon:        staticHTML["x-icon"],
			ModalIconBgColor: "bg-danger",
			ClearFormOnClose: false,
		}
		
		} else {
			modalProps = types.Modal{
				Type:             "success",
				ModalTitle:       "We received your request",
				ModalSubtitle:    "We will be in contact with you soon",
				ModalIcon:        staticHTML["check-icon"],
			ModalIconBgColor: "bg-success",
			ClearFormOnClose: true,
		}
	}
	
	return web.HTML(http.StatusOK, html, "components/modal.html", modalProps, nil)
}
func error404() *web.Response {
	return web.HTML(http.StatusOK, html, "pages/404.html", nil, nil)
}

func robotsTxt(r *http.Request) *web.Response {
	return web.Data(http.StatusOK, []byte(robotsContent), nil)
}