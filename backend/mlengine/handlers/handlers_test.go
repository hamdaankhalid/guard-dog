package handlers_test

import (
	"encoding/json"
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
		ModelsWithData:  []dal.Model{modelWithData},
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

// DeleteModel Tests

func TestDeleteModel(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	modelUuid := router.Queries.(*dal.MockQueries).Models[0].Id
	r, err := http.NewRequest("DELETE", "model/"+modelUuid.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelUuid.String()
	r = mux.SetURLVars(r, vars)
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
	model := router.Queries.(*dal.MockQueries).Models[0]
	// Make this models user not the user who will be making the call
	modelUuid := model.Id
	r, err := http.NewRequest("DELETE", "model/"+modelUuid.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelUuid.String()
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	user := middlewares.User{Id: model.UserId + 5} // User does not own the model
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

// TODO: GetMlNotification Tests

func TestGetMlNotification(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	expect := router.Queries.(*dal.MockQueries).MlNotifications[0]
	mlNotificationId := expect.Id

	r, err := http.NewRequest("GET", "ml-notification"+mlNotificationId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["mlNotificationId"] = mlNotificationId.String()
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 1}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetMlNotification(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
	dec := json.NewDecoder(w.Result().Body)
	var res dal.MlNotification
	err = dec.Decode(&res)
	if err != nil {
		t.Fail()
	}
	if res != expect {
		t.Fail()
	}
}

func TestGetMlNotificationNotYourMlNotification(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	expect := router.Queries.(*dal.MockQueries).MlNotifications[0]
	mlNotificationId := expect.Id

	r, err := http.NewRequest("GET", "ml-notification"+mlNotificationId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["mlNotificationId"] = mlNotificationId.String()
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 99}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetMlNotification(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusUnauthorized {
		t.Fail()
	}
}

func TestGetMlNotificationNotExists(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	mlNotificationId := uuid.New()

	r, err := http.NewRequest("GET", "ml-notification"+mlNotificationId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["mlNotificationId"] = mlNotificationId.String()
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	user := middlewares.User{Id: 99}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetMlNotification(w, r, user)

	// Assert
	if w.Result().StatusCode != http.StatusInternalServerError {
		t.Fail()
	}
}

// TODO: GetMlNotifications Tests

func TestGetMlNotifications(t *testing.T) {
}

func TestGetMlNotificationsWhenNoModelsExist(t *testing.T) {
}

// TODO: GetModel Tests

func TestGetModel(t *testing.T) {

}

func TestGetModelNotYourModel(t *testing.T) {

}

func TestGetModelNotExists(t *testing.T) {

}

// TODO: GetModels Tests
func TestGetModels(t *testing.T) {
}

func TestGetModelsWhenNoModelsExist(t *testing.T) {
}

// Health Handler Tests

func TestHealth(t *testing.T) {
	router := MockedPassingDependencyRouter()
	r, err := http.NewRequest("GET", "health", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()

	router.Health(w, r)

	if w.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}

// TODO: UploadModel Tests
func TestUploadModel(t *testing.T) {

}
