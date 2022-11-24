package main

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/database"
	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"
	"github.com/hamdaankhalid/mlengine/workers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func setup() error {
	godotenv.Load(".env")
	log.Println("Initializing DB")
	err := database.InitDb(false)
	return err
}

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.Health).Methods(http.MethodGet)

	router.Handle("/model", middlewares.NewAuth(handlers.UploadModel)).Methods(http.MethodPost)
	router.Handle("/model", middlewares.NewAuth(handlers.GetModels)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(handlers.GetModel)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(handlers.DeleteModel)).Methods(http.MethodDelete)
	router.Handle("/ml-notifications", middlewares.NewAuth(handlers.GetMlNotification)).Methods(http.MethodGet)

	log.Println("Listening And Serving ML Engine Up")
	log.Fatal(http.ListenAndServe(":6969", router))
}

func kickOffServer() error {

	listener, err := workers.NewListener()
	if err != nil {
		return err
	}

	go func() {
		err := listener.SubscribeAndConsume()
		if err != nil {
			log.Fatalf("Listen and Serve Exited: %s", err)
			return
		}
	}()

	// Blocking call so our above kicked off go routine will not exit early :)
	handleRequests()
	return nil
}

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = kickOffServer()
	if err != nil {
		log.Fatal(err)
	}
}
