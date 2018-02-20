package main

import (
	"log"
	"net/http"

	"github.com/bontequero/golang-test-assignment/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("can not load env variables: %v", err)
	}

	db, err := handlers.NewEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe("", router))
}
