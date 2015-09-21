package router

import (
	"gorestful/controller"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controller.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controller.TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		controller.TodoShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		controller.TodoCreate,
	},
	Route{
		"Users",
		"GET",
		"/users",
		controller.All,
	},
}
