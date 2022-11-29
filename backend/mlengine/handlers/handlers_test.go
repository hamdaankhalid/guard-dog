package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
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
	req, err := http.NewRequest("DELETE", "model/"+modelUuid.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelUuid.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: 1}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.DeleteModel(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}

func TestDeleteModelNotYourModel(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	model := router.Queries.(*dal.MockQueries).Models[0]
	// Make this models user not the user who will be making the call
	modelUuid := model.Id
	req, err := http.NewRequest("DELETE", "model/"+modelUuid.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelUuid.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: model.UserId + 5} // User does not own the model
	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.DeleteModel(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusUnauthorized {
		t.Fail()
	}
}

// GetMlNotification Tests

func TestGetMlNotification(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	expect := router.Queries.(*dal.MockQueries).MlNotifications[0]
	mlNotificationId := expect.Id

	req, err := http.NewRequest("GET", "ml-notification/"+mlNotificationId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["mlNotificationId"] = mlNotificationId.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: 1}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetMlNotification(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
	dec := json.NewDecoder(res.Result().Body)
	var result dal.MlNotification
	err = dec.Decode(&result)
	if err != nil {
		t.Fail()
	}
	if result != expect {
		t.Fail()
	}
}

func TestGetMlNotificationNotYourMlNotification(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	expect := router.Queries.(*dal.MockQueries).MlNotifications[0]
	mlNotificationId := expect.Id

	req, err := http.NewRequest("GET", "ml-notification/"+mlNotificationId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["mlNotificationId"] = mlNotificationId.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: 99}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetMlNotification(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusUnauthorized {
		t.Fail()
	}
}

func TestGetMlNotificationNotExists(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	mlNotificationId := uuid.New()

	req, err := http.NewRequest("GET", "ml-notification/"+mlNotificationId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["mlNotificationId"] = mlNotificationId.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: 99}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetMlNotification(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusInternalServerError {
		t.Fail()
	}
}

// GetMlNotifications Tests

func TestGetMlNotifications(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	userId := router.Queries.(*dal.MockQueries).MlNotifications[0].UserId
	expected := router.Queries.(*dal.MockQueries).MlNotifications
	req, err := http.NewRequest("GET", "ml-notification", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	user := middlewares.User{Id: userId}

	// Invoke
	router.GetMlNotifications(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
	dec := json.NewDecoder(res.Result().Body)
	var result []dal.MlNotification
	err = dec.Decode(&result)
	if err != nil {
		t.Fail()
	}
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

// GetModel Tests

func TestGetModel(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	expect := router.Queries.(*dal.MockQueries).ModelsWithData[0]
	modelId := expect.Id

	req, err := http.NewRequest("GET", "model/"+modelId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelId.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: expect.UserId}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetModel(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
	dec := json.NewDecoder(res.Result().Body)
	var result dal.Model
	err = dec.Decode(&result)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(result, expect) {
		t.Fail()
	}
}

func TestGetModelNotYourModel(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	expect := router.Queries.(*dal.MockQueries).ModelsWithData[0]
	modelId := expect.Id

	req, err := http.NewRequest("GET", "model/"+modelId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelId.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: expect.UserId + 5}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetModel(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusUnauthorized {
		t.Fail()
	}
}

func TestGetModelNotExists(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	modelId := uuid.New()

	req, err := http.NewRequest("GET", "model/"+modelId.String(), nil)
	if err != nil {
		t.Fatal(err)
	}
	vars := make(map[string]string)
	vars["modelId"] = modelId.String()
	req = mux.SetURLVars(req, vars)
	res := httptest.NewRecorder()
	user := middlewares.User{Id: 99}

	if err != nil {
		t.Fatal(err)
	}

	// Invoke
	router.GetModel(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusInternalServerError {
		t.Fail()
	}
}

// GetModels Tests

func TestGetModels(t *testing.T) {
	// Setup
	router := MockedPassingDependencyRouter()
	userId := router.Queries.(*dal.MockQueries).Models[0].UserId
	expected := router.Queries.(*dal.MockQueries).Models
	req, err := http.NewRequest("GET", "model", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	user := middlewares.User{Id: userId}

	// Invoke
	router.GetModels(res, req, user)

	// Assert
	if res.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
	dec := json.NewDecoder(res.Result().Body)
	var result []dal.ModelWithoutData
	err = dec.Decode(&result)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

// Health Handler Tests

func TestHealth(t *testing.T) {
	router := MockedPassingDependencyRouter()
	req, err := http.NewRequest("GET", "health", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()

	router.Health(res, req)

	if res.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}

// UploadModel Tests

func TestUploadModel(t *testing.T) {
	// setup
	router := MockedPassingDependencyRouter()
	filePath := "../test-resources/fakemodel.onnx"
	fieldName := "model"
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)

	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}

	w, err := mw.CreateFormFile(fieldName, filePath)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := io.Copy(w, file); err != nil {
		t.Fatal(err)
	}

	// close the writer before making the request
	mw.Close()

	user := middlewares.User{Id: 1}
	req := httptest.NewRequest(http.MethodPost, "/upload", body)

	req.Header.Add("Content-Type", mw.FormDataContentType())

	res := httptest.NewRecorder()

	// invoke
	router.UploadModel(res, req, user)

	// assert
	if res.Result().StatusCode != http.StatusCreated {
		t.Fail()
	}
}
