package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/middlewares"
)

func (router *Router) DeleteModel(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	vars := mux.Vars(r)
	userId := user.Id
	modelId, err := uuid.Parse(vars["modelId"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model, err := dal.RetrieveModel(modelId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if model.UserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = dal.DeleteModel(modelId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
