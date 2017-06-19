package main

import (
	"log"
	"net/http"
)

var signIn = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	log.Fatal("Not implemented")
})
