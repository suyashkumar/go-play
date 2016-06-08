/*
	The bind package contains a function to create a new httprouter Router and bind handlers
	to appropiate routes

	@author: Suyash Kumar <suyashkumar2003@gmail.com>
*/
package bind

import (
	"github.com/julienschmidt/httprouter"
	"routes/home"
	"routes/timeroute"
)

// BindRoutes creates the routetable by binding routes to the proper handlers in a new
// httprouter object. This function returns a pointer to that Router object.
func BindRoutes() *httprouter.Router {
	// Get a router
	router := httprouter.New()

	// Begin route table:
	router.GET("/", home.Index)
	router.GET("/time", timeroute.CurrTimeANSIC)
	router.GET("/time/:name", timeroute.CurrTimeANSICName)

	return router

}
