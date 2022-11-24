package workers

import (
	"fmt"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type KafkaMsgProcessor func(*kafka.Message)

func RouteTask(key string) KafkaMsgProcessor {
	switch key {
	case "video_upload":
		return RunInference
	default:
		return func(_ *kafka.Message) {}
	}
}

// TODO:
func RunInference(msg *kafka.Message) {
	fmt.Println(msg)
}
