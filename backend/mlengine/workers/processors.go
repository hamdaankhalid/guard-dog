package workers

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type VideoUploadEvent struct {
	Url        string `json:"url"`
	UserId     int    `json:"user_id"`
	DeviceName string `json:"deviceName"`
	SessionId  int    `json:"sessionId"`
	Part       int    `json:"part"`
}

type KafkaMsgProcessor func(*kafka.Message)

func RouteTask(topic *string) KafkaMsgProcessor {
	switch *topic {
	case "video-upload":
		return RunInference
	default:
		return func(_ *kafka.Message) {}
	}
}

// TODO:
func RunInference(msg *kafka.Message) {
	var event VideoUploadEvent
	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(event.UserId)
	fmt.Println(event.Url)
}
