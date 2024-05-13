package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

var (
	//go:embed all:templates/**
	templateFS embed.FS
	//parsed templates
	html *template.Template
)

func main() {

	//parse templates and create relations so that templates can reference eachother
	var err error
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
