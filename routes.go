package main

import (
	"net/http"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

func index(r *http.Request) *web.Response {
	return web.HTML(http.StatusOK, html, "pages/index.html", nil, nil)
}
