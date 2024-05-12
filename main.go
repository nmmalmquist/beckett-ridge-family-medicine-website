package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

var (

	//go:embed css/output.css
	css embed.FS
)

func main() {

	// Add Routes
	router := http.NewServeMux()
	router.Handle("/", web.Action(index))
	// router.Handle("/css/output.css", http.FileServer(http.FS(css)))
	router.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// Logging and tracing
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := logging(logger)(router)

	ADDRESS := "localhost:8000"
	fmt.Println("Started web server on", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, middleware))
}
