package main

import (
	"net/http"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/types"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)


func index(r *http.Request) *web.Response {
	type IndexPageData struct {
		CardFacts []types.CardFact
	}
	
	data := IndexPageData{
        CardFacts: []types.CardFact{
            {Title: staticHTML["dr-icon"]},
        },
    }
	return web.HTML(http.StatusOK, html, "pages/index.html", data, nil)
}
