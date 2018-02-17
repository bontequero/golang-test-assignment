package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bontequero/golang-test-assignment/db"
	"github.com/bontequero/golang-test-assignment/handlers"

	"github.com/joho/godotenv"
)

const (
	postgresUrl = "POSTGRES_URL"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("can not load env variables: %v", err)
	}
	db, err := db.NewDB(os.Getenv(postgresUrl))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe("", router))
}
