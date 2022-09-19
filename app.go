package main

import (
	"log"
	"net/http"
	"portfolio/backend"
	"portfolio/frontend"
)

func main() {
	if err := backend.Minifier("assets/js/compile.js", "text/javascript", "assets/js/components"); err != nil {
		log.Fatal(err)
	}
	if err := backend.Minifier("assets/styles/style.css", "text/css", "assets/styles"); err != nil {
		log.Fatal(err)
	}

	mux := frontend.Routes("assets")

	log.Fatal(http.ListenAndServe(":9001", mux))
}
