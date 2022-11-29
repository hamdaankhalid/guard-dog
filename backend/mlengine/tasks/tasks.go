package tasks

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/simple"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const HUMAN_DETECTION_MODEL = "human-detection-model.onnx"

const UploadModelTaskName = "uploadModelTask"

type VideoUploadEvent struct {
	Url        string `json:"url"`
	UserId     int    `json:"user_id"`
	DeviceName string `json:"deviceName"`
	SessionId  int    `json:"sessionId"`
	Part       int    `json:"part"`
}

const InferenceOnModelTaskName = "inferenceOnModelTask"

type UploadModelReq struct {
	File    *multipart.File
	Handler *multipart.FileHeader
	UserId  int
}

func ParallelModelInferenceTask(msg *kafka.Message, queries dal.IQueries) {
	var event VideoUploadEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Println(err)
	}
	userId := event.UserId
	models, err := queries.RetrieveAllModels(userId)
	if err != nil {
		log.Printf("Error getting models for userId: %d", userId)
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(models))
	for _, model := range models {
		go inferenceOnModel(&model, &event, &wg, queries)
	}
	wg.Wait()
}

func inferenceOnModel(model *dal.ModelWithoutData, event *VideoUploadEvent, wg *sync.WaitGroup, queries dal.IQueries) {
	switch model.Filename {
	case HUMAN_DETECTION_MODEL:
		humanDetection(model, event, queries)
	default:
		log.Printf("Unregistered Model: %s", model.Filename)
	}
}

func humanDetection(model *dal.ModelWithoutData, event *VideoUploadEvent, queries dal.IQueries) {
	modelWithData, err := queries.RetrieveModel(model.Id)
	if err != nil {
		log.Printf("Could not pull model data for mode Id: %s", model.Id)
		return
	}
	backend := simple.NewSimpleGraph()
	loadedModel := onnx.NewModel(backend)
	err = loadedModel.UnmarshalBinary(modelWithData.ModelFile)
	if err != nil {
		log.Printf("Error Marshalling Model: %s, Err: %s", modelWithData.Filename, err)
		return
	}
	videoData, err := http.Get(event.Url)
	if err != nil {
		log.Printf("Error Retrieving From URL: %s", event.Url)
		return
	}
	defer videoData.Body.Close()
	log.Println("videodata", videoData)
}

func UploadModelTask(uploadModelReq *UploadModelReq, queries dal.IQueries) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
		return
	}

	bytes := bytes.NewBuffer(nil)
	n, err := io.Copy(bytes, *uploadModelReq.File)
	if err != nil || n == 0 {
		log.Println("Read Error", err)
		return
	}

	model := dal.Model{ModelFile: bytes.Bytes(), Id: uuid, Filename: uploadModelReq.Handler.Filename, UserId: uploadModelReq.UserId}
	err = queries.UploadModel(&model)
	if err != nil {
		log.Println("Error uploading model file: ", model.Filename)
	}
}
