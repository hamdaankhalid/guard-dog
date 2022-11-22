package main

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/database"
	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"

	"github.com/joho/godotenv"
)

func setup() {
	godotenv.Load(".env")
	database.InitDb()
}

func handlerRequests() {
	http.HandleFunc("/health", handlers.Health)

	http.Handle("/model", middlewares.NewAuth(handlers.UploadModel))

	log.Println("Booting up ML Engine Up")
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func main() {
	setup()
	handlerRequests()
}
