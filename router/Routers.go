package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorestful/controller"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = controller.Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path("/v1" + route.Pattern).
			Name(route.Name).
			Handler(handler)
		fmt.Println("/v1" + route.Pattern)
	}

	return router
}
