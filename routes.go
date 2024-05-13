package main

import (
	"net/http"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/types"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)


func index(r *http.Request) *web.Response {
	type IndexPageData struct {
		CardFacts []types.CardFact
		TabFacts []types.TabFact
	}
	
	data := IndexPageData{
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
    }
	return web.HTML(http.StatusOK, html, "pages/index.html", data, nil)
}
