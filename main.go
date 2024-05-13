package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

var (
	//go:embed all:templates/**
	templateFS embed.FS
	//parsed templates
	html *template.Template
	//go:embed all:static_html/**
	staticHTMLFS embed.FS
	// map of static html components
	staticHTML map[string]string
)

func main() {
	var err error
	// read in html for static components it map for use in data for templates
	staticHTML = make(map[string]string)
	err = fs.WalkDir(staticHTMLFS, "static_html", func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir(){
			
			if err != nil {
				return err
			}
			
			data, e :=  os.ReadFile(path)
			if e != nil {
				return e
			}
			
			parts:=strings.Split(path, string(os.PathSeparator))
			name := strings.Split(parts[len(parts)-1], ".")[0]
			
			staticHTML[name] = string(data)
		}
		return nil
	})

	//parse templates and create relations so that templates can reference eachother
	if err != nil {
		panic(err)
	}
	html, err = web.TemplateParseFSRecursive(templateFS, ".html", true, nil)
	if err != nil {
		panic(err)
	}

	// Add Routes
	router := http.NewServeMux()
	// Allows access to images, css, and js files
	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.Handle("/", web.Action(index))

	// Logging and tracing
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := logging(logger)(router)

	ADDRESS := "localhost:8000"
	fmt.Println("Started web server on", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, middleware))

}
