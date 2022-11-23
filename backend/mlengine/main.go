package main

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"

	"github.com/joho/godotenv"
)

func setup() error {
	godotenv.Load(".env")
	log.Println("Initializing DB")
	// err := database.InitDb()
	return nil
}

func handlerRequests() {
	http.HandleFunc("/health", handlers.Health)

	http.Handle("/model", middlewares.NewAuth(handlers.UploadModel))

	log.Println("Booting up ML Engine Up")
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
		return
	}
	handlerRequests()
}
