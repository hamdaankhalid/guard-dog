package main

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func setup() error {
	godotenv.Load(".env")
	log.Println("Initializing DB")
	// err := database.InitDb()
	return nil
}

func handlerRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.Health).Methods(http.MethodGet)

	router.Handle("/model", middlewares.NewAuth(handlers.UploadModel)).Methods(http.MethodPost)
	router.Handle("/model", middlewares.NewAuth(handlers.GetModels)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(handlers.GetModel)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(handlers.GetModel)).Methods(http.MethodDelete)

	log.Println("Booting up ML Engine Up")
	log.Fatal(http.ListenAndServe(":6969", router))
}

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
		return
	}
	handlerRequests()
}
