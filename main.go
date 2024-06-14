package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS
	//parsed templates
	html *template.Template
	//go:embed all:static/*
	staticFS embed.FS
	//go:embed robots.txt
	robotsContent string

	// Services
	appServices *AppServices
)

func main() {
	godotenv.Load()

	var err error

	// Specify custom functions to use inside html templates
	funcMap := template.FuncMap{
		"isEven": func(i int) bool {
			return i%2 == 0
		},
	}
	html, err = web.TemplateParseFSRecursive(templateFS, ".html", true, funcMap)
	if err != nil {
		panic(err)
	}

	//Add Services
	appServices = InitServices()

	// Add Routes
	router := http.NewServeMux()
	// Allows access to images, css, and js files
	router.Handle("/static/", http.FileServer(http.FS(staticFS)))
	// Pages
	router.Handle("/", web.Action(index))
	router.Handle("/providers", web.Action(providers))
	router.Handle("/request-appointment", web.Action(requestAppointment))
	router.Handle("/privacy-policy", web.Action(privacyPolicy))
	// API routes
	router.Handle("/api/request-appointment", web.Action(requestAppointmentPOST))

	// Utility pages
	router.Handle("/robots.txt", web.Action(robotsTxt))
	router.Handle("/error", web.Action(errorPage))

	// Logging and tracing
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := logging(logger)(router)

	ADDRESS := "0.0.0.0:8000"
	fmt.Println("Started web server on", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, middleware))
}
