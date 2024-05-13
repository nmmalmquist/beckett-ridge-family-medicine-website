package web

import (
	"bytes"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
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
func HTML(status int, t *template.Template, templateName string, data interface{}, headers Headers) *Response {
	//render template to buffer
	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, templateName, data); err != nil {
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

// TemplateParseFSRecursive recursively parses all templates in the FS with the given extension.
// File paths are used as template names to support duplicate file names.
// Use nonRootTemplateNames to exclude root directory from template names
// (e.g. index.html instead of templates/index.html)
func TemplateParseFSRecursive(
	templates fs.FS,
	ext string,
	nonRootTemplateNames bool,
	funcMap template.FuncMap) (*template.Template, error) {

	root := template.New("")
	err := fs.WalkDir(templates, "templates", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && strings.HasSuffix(path, ext) {
			if err != nil {
				return err
			}
			b, err := fs.ReadFile(templates, path)
			if err != nil {
				return err
			}
			name := ""
			if nonRootTemplateNames {
				//name the template based on the file path (excluding the root)
				parts := strings.Split(path, string(os.PathSeparator))
				name = strings.Join(parts[1:], string(os.PathSeparator))
			}
			t := root.New(name).Funcs(funcMap)
			_, err = t.Parse(string(b))
			if err != nil {
				return err
			}
		}
		return nil
	})
	return root, err
}
