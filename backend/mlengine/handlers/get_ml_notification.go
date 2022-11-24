package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/middlewares"
)

func GetMlNotification(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	userId := user.Id

	mlNotifications, err := dal.RetrieveAllMlNotifications(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(mlNotifications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
