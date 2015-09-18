package main

import (
	"gorestful/router"
	"log"
	"net/http"
)

func main() {

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
