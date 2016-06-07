package home

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Route handlers:
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hey there")

}
