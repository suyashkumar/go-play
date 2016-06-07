package main

import (
	"bind"
	"log"
	"net/http"
)

func main() {
	// Init Webserver
	router := bind.BindRoutes() // bind routes based on route table
	log.Fatal(http.ListenAndServe(":9002", router))
}
