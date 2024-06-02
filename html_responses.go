package main

import (
	"net/http"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/types"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

func GetErrorModal(status int) *web.Response{
	modalProps := types.Modal{
		Type:             "error",
		ModalTitle:       "An error occured",
		ModalSubtitle:    "Please try again later",
		ModalIcon:        staticHTML["x-icon"],
		ModalIconBgColor: "bg-danger",
		ClearFormOnClose: false,
	}
	return web.HTML(http.StatusOK, html, "components/modal.html", modalProps, nil)
}

func GetSuccessModal() *web.Response{
	modalProps := types.Modal{
		Type:             "success",
		ModalTitle:       "We received your request",
		ModalSubtitle:    "We will be in contact with you soon",
		ModalIcon:        staticHTML["check-icon"],
		ModalIconBgColor: "bg-success",
		ClearFormOnClose: true,
	}
	return web.HTML(http.StatusOK, html, "components/modal.html", modalProps, nil)
}