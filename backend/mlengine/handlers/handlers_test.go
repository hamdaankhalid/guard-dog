package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"
	"github.com/hamdaankhalid/mlengine/processingqueue"
)

func MockedPassingDependencyRouter() *handlers.Router {
	id1 := uuid.New()
	id2 := uuid.New()

	model1 := dal.ModelWithoutData{
		Id:       id1,
		Filename: "f1.onnx",
		UserId:   1,
	}

	model2 := dal.ModelWithoutData{
		Id:       id2,
		Filename: "f2.onnx",
		UserId:   1,
	}

	modelWithData := dal.Model{
		Id:        id2,
		Filename:  "f2.onnx",
		UserId:    1,
		ModelFile: []byte{},
	}

	id3 := uuid.New()
	modelId1 := uuid.New()
	id4 := uuid.New()
	modelId2 := uuid.New()

	notification1 := dal.MlNotification{
		Id:         id3,
		DeviceName: "dev1",
		SessionId:  1,
		Part:       1,
		ModelId:    modelId1,
		UserId:     1}

	notification2 := dal.MlNotification{
		Id:         id4,
		DeviceName: "dev1",
		SessionId:  2,
		Part:       1,
		ModelId:    modelId2,
		UserId:     1}

	queries := dal.MockQueries{
		ErrorOnly:       false,
		Error:           nil,
		MlNotifications: []dal.MlNotification{notification1, notification2},
		Models:          []dal.ModelWithoutData{model1, model2},
		Model:           modelWithData,
	}
	testQueue := &processingqueue.MockQueue{InnerState: []string{}}
	router := handlers.NewRouter(testQueue, &queries)
	return router
}

func MockedFailingDependencyRouter(err error) *handlers.Router {
	queries := dal.MockQueries{ErrorOnly: true, Error: err}
	testQueue := &processingqueue.MockQueue{InnerState: []string{}}
	router := handlers.NewRouter(testQueue, &queries)
	return router
}

// Delete Model Tests

func TestDeleteModel(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	modelUuid := router.Queries.(*dal.MockQueries).Model.Id
	r, err := http.NewRequest("DELETE", "model"+modelUuid.String(), nil)
	vars := make(map[string]string)
	vars["modelId"] = modelUuid.String()
	r = mux.SetURLVars(r, vars)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 1}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.DeleteModel(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}

func TestDeleteModelNotYourModel(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	modelUuid := router.Queries.(*dal.MockQueries).Model.Id
	// Make this models user not the user who will be making the call
	router.Queries.(*dal.MockQueries).Model.UserId = 2
	r, err := http.NewRequest("DELETE", "model"+modelUuid.String(), nil)
	vars := make(map[string]string)
	vars["modelId"] = modelUuid.String()
	r = mux.SetURLVars(r, vars)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 1}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.DeleteModel(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusUnauthorized {
		t.Fail()
	}
}

func TestDeleteModelInvalidUuid(t *testing.T) {
	// Setup
	r, err := http.NewRequest("DELETE", "model", nil)
	vars := make(map[string]string)
	vars["modelId"] = "notuuid"
	r = mux.SetURLVars(r, vars)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 1}
	router := MockedPassingDependencyRouter()

	// Invoke
	router.DeleteModel(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Fail()
	}
}
