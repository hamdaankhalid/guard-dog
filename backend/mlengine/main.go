package main

import (
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/handlers"
)

func handlerRequests() {
	http.HandleFunc("/health", handlers.Health)
	http.HandleFunc("/model", handlers.UploadModel)

	log.Println("Booting up ML Engine Up")
	log.Fatal(http.ListenAndServe(":6969", nil))
}

func main() {
	handlerRequests()
}
