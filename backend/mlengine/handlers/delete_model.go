package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hamdaankhalid/mlengine/middlewares"
)

func (router *Router) DeleteModel(w http.ResponseWriter, r *http.Request, user middlewares.User) {
	vars := mux.Vars(r)
	userId := user.Id
	modelUuid, ok := vars["modelId"]
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	modelId, err := uuid.Parse(modelUuid)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	model, err := router.Queries.RetrieveModel(modelId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if model.UserId != userId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = router.Queries.DeleteModel(modelId)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
