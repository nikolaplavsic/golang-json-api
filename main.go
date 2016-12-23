package main

import (
	"log"
	"net/http"
)

func init() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
