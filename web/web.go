package web

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"text/template"
)

type Headers map[string]string
type Response struct {
	// StatusCode
	Status int
	// Content Type to writer
	ContentType string
	// Content to be written to the response writer
	Content io.Reader
	// Headers to be written to the response writer
	Headers Headers
}

// Write writes a response to an http.ResponseWriter
func (response *Response) Write(rw http.ResponseWriter) {
	if response != nil {
		if response.ContentType != "" {
			rw.Header().Set("Content-Type", response.ContentType)
		}
		for k, v := range response.Headers {
			rw.Header().Set(k, v)
		}
		rw.WriteHeader(response.Status)
		_, err := io.Copy(rw, response.Content)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}

type Action func(r *http.Request) *Response

// Action's http.Handler implementation
func (a Action) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := a(r)
	response.Write(rw)
}

// Data returns a data response
func Data(status int, content []byte, headers Headers) *Response {
	return &Response{
		Status:  status,
		Content: bytes.NewBuffer(content),
		Headers: headers,
	}
}

// Empty returns an empty http response
func EmptyHTTP(status int) *Response {
	return Data(status, []byte(""), nil)
}

// HTML renders an html template to a web response
func HTML(status int, templateFilePath string, data interface{}, headers Headers) *Response {
	//render template to buffer
	var buf bytes.Buffer
	t := template.Must(template.ParseFiles(templateFilePath))
	if err := t.Execute(&buf, data); err != nil {
		log.Println(err)
		return EmptyHTTP(http.StatusInternalServerError)
	}
	return &Response{
		Status:      status,
		ContentType: "text/html",
		Content:     &buf,
		Headers:     headers,
	}
}
