package main

import (
	"portfolio/backend"
)

func main() {
	if err := backend.Minifier("assets/js/compile.js", "text/javascript", "assets/js/components"); err != nil {
		panic(err.Error())
	}
	if err := backend.Minifier("assets/styles/style.css", "text/css", "assets/styles"); err != nil {
		panic(err.Error())
	}

	//mux := frontend.Routes("assets")
	//
	//log.Fatal(http.ListenAndServe(":9001", mux))
}
