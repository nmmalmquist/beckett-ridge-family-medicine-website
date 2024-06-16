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
	type ProvidersPageData struct {
		ActivePage string
	}
	data := ProvidersPageData{
		ActivePage: "providers",
	}
	return web.HTML(http.StatusOK, html, "pages/providers.html", data, nil)
}
func weightManagement(r *http.Request) *web.Response {
	type WeightManagementPageData struct {
		ActivePage string
		CardFacts  []types.CardFact
	}

	data := WeightManagementPageData{
		ActivePage: "weight-management",
		CardFacts: []types.CardFact{
			{Icon: GetHTMLFromRootTemplate("components/lightning-icon.html"), Title: "Understand Your Health", Description: "Feeling tired and unmotivated can make it challenging to be productive, but our services are designed to help you boost your daily energy levels and achieve weight loss. Our comprehensive weight loss care program includes personalized meal plans that provide balanced nutrition, ensuring you stay energized throughout the day. We offer regular exercise routines tailored to your fitness level, which not only aid in weight loss but also enhance your vitality. Additionally, our stress management techniques, such as guided meditation and breathing exercises, help reduce fatigue and improve overall well-being. By utilizing our expert guidance and support, you can achieve your weight loss goals and enjoy a more energized, productive life."},
			{Icon: GetHTMLFromRootTemplate("components/check-icon.html"), Title: "Confidence", Description: "Gaining confidence in our weight loss program is essential for your success, and we ensure that every step is tailored to your unique needs. Our program is backed by scientific research and has helped countless individuals achieve their weight loss goals effectively and sustainably. With personalized support from our experienced professionals, you can trust that you are in capable hands. Regular progress tracking and adjustments ensure that the program adapts to your evolving needs, maximizing your results. By seeing consistent improvements and receiving ongoing encouragement, you'll gain the confidence that our weight loss program truly works."},
			{Icon: GetHTMLFromRootTemplate("components/vitality-icon.html"), Title: "Reduced Body Fat", Description: "Our program's proven results in decreasing body fat demonstrate its effectiveness and reliability. Countless individuals have successfully reduced their body fat percentage through our tailored approach, showcasing the program's ability to deliver tangible outcomes. Our scientific methodology, combined with personalized support from experienced professionals, ensures you achieve the best possible results. Consistent progress tracking and adaptive strategies further validate the success of our program in helping you reach your goals. Seeing these real, measurable improvements will give you the confidence that our program truly works."},
		},
	}
	return web.HTML(http.StatusOK, html, "pages/weight-management.html", data, nil)
}
func privacyPolicy(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "pages/privacy-policy.html", nil, nil)
}
func requestAppointment(r *http.Request) *web.Response {
	type RequestAppointmentPageData struct {
		ActivePage string
	}
	data := RequestAppointmentPageData{
		ActivePage: "request-appointment",
	}
	return web.HTML(http.StatusOK, html, "pages/request-appointment.html", data, nil)
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
