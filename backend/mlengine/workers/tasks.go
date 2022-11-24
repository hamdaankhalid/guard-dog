package workers

import "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"

type KafkaMsgProcessor func(*kafka.Message)

func RouteTask(key string) KafkaMsgProcessor {
	switch key {
	case "video_upload":
		return RunInference
	default:
		return func(_ *kafka.Message) {}
	}
}

// TODO
func RunInference(_ *kafka.Message) {

}
