package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hamdaankhalid/mlengine/middlewares"
)

func (router *Router) GetModels(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	userId := user.Id

	models, err := router.Queries.RetrieveAllModels(userId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(models)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)
}
