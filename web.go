/*
	The main entrypoint of this application. The purpose of this application is to build a
	basic webservice and learn about Go.

	@author: Suyash Kumar <suyashkumar2003@gmail.com>
*/
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
