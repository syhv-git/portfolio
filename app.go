package main

import (
	"go-networking/frontend"
	"log"
	"net/http"
)

func main() {
	mux := frontend.Routes("assets")

	log.Fatal(http.ListenAndServe(":9001", mux))
}
