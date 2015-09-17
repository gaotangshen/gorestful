package main

import (
	"log"
	"net/http"
	"test/router"
)

func main() {

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
