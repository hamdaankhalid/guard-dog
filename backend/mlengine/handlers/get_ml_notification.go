package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hamdaankhalid/mlengine/middlewares"
)

func (router *Router) GetMlNotification(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	vars := mux.Vars(r)
	userId := user.Id
	mlNotificationId, err := uuid.Parse(vars["mlNotificationId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	mlNotification, err := router.Queries.RetrieveMlNotification(mlNotificationId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if mlNotification.UserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(mlNotification)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
