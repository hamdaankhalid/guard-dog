package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hamdaankhalid/mlengine/middlewares"
	"github.com/hamdaankhalid/mlengine/processingqueue"
)

type Router struct {
	processingQueue processingqueue.IQueue
	Routing         *mux.Router
}

func NewRouter(processingQueue processingqueue.IQueue) *Router {
	router := mux.NewRouter()
	res := &Router{processingQueue: processingQueue, Routing: router}

	router.HandleFunc("/health", res.Health).Methods(http.MethodGet)
	router.Handle("/model", middlewares.NewAuth(res.UploadModel)).Methods(http.MethodPost)
	router.Handle("/model", middlewares.NewAuth(res.GetModels)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(res.GetModel)).Methods(http.MethodGet)
	router.Handle("/model/{modelId}", middlewares.NewAuth(res.DeleteModel)).Methods(http.MethodDelete)
	router.Handle("/ml-notification", middlewares.NewAuth(res.GetMlNotifications)).Methods(http.MethodGet)
	router.Handle("/ml-notification/{mlNotificationId}", middlewares.NewAuth(res.GetMlNotification)).Methods(http.MethodGet)

	return res
}
