package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.Health).Methods(http.MethodGet)
	router.Handle("/model", middlewares.NewAuth(handlers.UploadModel)).Methods(http.MethodPost)
	router.Handle("/model", middlewares.NewAuth(handlers.GetModels)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(handlers.GetModel)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(handlers.DeleteModel)).Methods(http.MethodDelete)
	router.Handle("/ml-notification", middlewares.NewAuth(handlers.GetMlNotifications)).Methods(http.MethodGet)
	router.Handle("/ml-notification/{mlNotificationId}", middlewares.NewAuth(handlers.GetMlNotification)).Methods(http.MethodGet)

	return router
}
