package processingqueue

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/owulveryck/onnx-go"
	"github.com/owulveryck/onnx-go/backend/simple"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const HUMAN_DETECTION_MODEL = "human-detection-model.onnx"

func parallelModelInferenceTask(msg *kafka.Message, queries *dal.Queries) {
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

func inferenceOnModel(model *dal.ModelWithoutData, event *VideoUploadEvent, wg *sync.WaitGroup, queries *dal.Queries) {
	switch model.Filename {
	case HUMAN_DETECTION_MODEL:
		humanDetection(model, event, queries)
	default:
		log.Printf("Unregistered Model: %s", model.Filename)
	}
}

func humanDetection(model *dal.ModelWithoutData, event *VideoUploadEvent, queries *dal.Queries) {
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
