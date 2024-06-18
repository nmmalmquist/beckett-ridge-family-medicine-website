package main

import (
	"bytes"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

func GetHTMLFromRootTemplate(name string) *bytes.Buffer {
	byt, err := web.GetHTMLFromTemplate(html, name)
	if err != nil {
		return new(bytes.Buffer)
	}

	return &byt
}
