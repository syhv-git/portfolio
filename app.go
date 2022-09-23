package main

import (
	"github.com/syhv-git/minifier"
	"log"
	"net/http"
	"portfolio/frontend"
)

func main() {
	if err := minifier.Minifier("assets/js/compile.js", "text/javascript", "assets/js/components"); err != nil {
		log.Fatal(err)
	}
	if err := minifier.Minifier("assets/styles/style.css", "text/css", "assets/styles/navfoot.css", "assets/styles/landing.css", "assets/styles/responsive.css"); err != nil {
		log.Fatal(err)
	}

	mux := frontend.Routes("assets")

	log.Fatal(http.ListenAndServe(":9001", mux))
}
