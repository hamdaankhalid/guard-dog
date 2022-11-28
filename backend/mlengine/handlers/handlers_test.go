package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/middlewares"
	"github.com/hamdaankhalid/mlengine/processingqueue"
)

func MockedPassingDependencyRouter() *handlers.Router {
	queries := dal.MockQueries{ErrorOnly: false, Error: nil}
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

func TestDeleteModelInvalidUuid(t *testing.T) {
	// Setup
	r, err := http.NewRequest("DELETE", "/model/notavaliduuid", nil)
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

// TODO: DAL DependencyInjection
func TestDeleteModelNotYourModel(t *testing.T) {
	// Setup
	uuid, err := uuid.NewUUID()
	if err != nil {
		t.Fatal(err)
	}

	// mock dal.RetrieveModel & dal.DeleteModel
	r, err := http.NewRequest("DELETE", "/model/"+uuid.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 1}
	router, err := mockedDependencyRouter()
	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.DeleteModel(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Fail()
	}
}
