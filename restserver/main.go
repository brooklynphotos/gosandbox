package main

import (
	"log"
	"net/http"
)

func main() {
	router := MakeRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
