package main

import (
	"fmt"
	"net/http"
	"os"
	"restexample/api/v1/handler"
	"restexample/config"
	"restexample/db"
	"restexample/logadapter"

	"github.com/joho/godotenv"
)

func main() {

	boot()

	if !config.Mock() {
		defer db.Close()
	}

	mux := http.NewServeMux()

	mux.Handle("/ping", &handler.PingHandler{})
	mux.Handle("/actors/", &handler.ActorHandler{})

	prt := fmt.Sprintf(":%d", config.Port())

	logadapter.Log.Infof("Starting mode %s server on %s", config.Environment(), prt)
	err := http.ListenAndServe(prt, mux)

	if err != nil {
		logadapter.Log.Fatal(err)
	}
}

func boot() {
	_ = godotenv.Load()

	config.LoadCliConfig()

	if config.Environment() == "development" {

		err := db.Open(os.Getenv("DEV_API_DB_DSN"), config.Debug(), config.Mock())

		if err != nil {
			logadapter.Log.Fatal("Couldnt open database")
		}

	} else if config.Environment() == "production" {

		err := db.Open(os.Getenv("PROD_API_DB_DSN"), config.Debug(), config.Mock())

		if err != nil {
			logadapter.Log.Fatal("Couldnt open database")
		}
	} else {
		logadapter.Log.Fatal("Invalid environment")
	}
}
