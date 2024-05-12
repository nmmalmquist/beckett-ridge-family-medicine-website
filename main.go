package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"nickmalmquist.com/beckett-ridge-family-medicine-website/web"
)

func main() {

	// Add Routes
	router := http.NewServeMux()
	router.Handle("/", web.Action(index))

	// Logging and tracing
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	middleware := logging(logger)(router)

	ADDRESS := "localhost:8000"
	fmt.Println("Started web server on", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, middleware))
}
