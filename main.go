package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/email"
	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

var (
	//go:embed all:templates/*
	templateFS embed.FS
	//parsed templates
	html *template.Template
	//go:embed all:static_html/*
	staticHTMLFS embed.FS
	// map of static html components
	staticHTML map[string]string
	//go:embed all:static/*
	staticFS embed.FS

	// Services
	emailService *email.EmailService
)

func main() {
	godotenv.Load()

	var err error
	// read in html for static components it map for use in data for templates
	staticHTML = make(map[string]string)
	err = parseStaticHtml(staticHTMLFS)

	//parse templates and create relations so that templates can reference eachother
	if err != nil {
		panic(err)
	}
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
	initServices()

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

	// Logging and tracing
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := logging(logger)(router)

	ADDRESS := "0.0.0.0:8000"
	fmt.Println("Started web server on", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, middleware))

}

func initServices() {
	emailService = email.CreateEmailService()
}
