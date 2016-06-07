package bind

import (
	"github.com/julienschmidt/httprouter"
	"routes/home"
)

func BindRoutes() *httprouter.Router {
	router := httprouter.New()

	// Begin route table:
	router.GET("/", home.Index)

	return router

}
