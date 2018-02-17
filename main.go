package main

import (
	"log"
	"net/http"

	"github.com/bontequero/golang-test-assignment/handlers"
)

func main()  {
	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe("", router))
}