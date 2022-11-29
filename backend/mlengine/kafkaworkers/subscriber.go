package kafkaworkers

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hamdaankhalid/mlengine/processingqueue"
	"github.com/hamdaankhalid/mlengine/tasks"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

/**
	Listener is created via NewListener, it runs a blocking method
	which listens to incoming topic messages and processes them in
	it's own thread.
**/
type Listener struct {
	consumer        *kafka.Consumer
	topics          []string
	processingQueue processingqueue.IQueue
}

func NewListener(processingQueue processingqueue.IQueue) (*Listener, error) {
	topics := []string{"video-upload"}

	consumer, err := initConsumer()
	if err != nil {
		return nil, err
	}
	return &Listener{consumer: consumer, topics: topics, processingQueue: processingQueue}, nil
}

func initConsumer() (*kafka.Consumer, error) {
	kafkaServers := os.Getenv("KAFKA_SERVERS")
	kafkaGroupId := os.Getenv("KAFKA_GROUP_ID")

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaServers,
		"group.id":          kafkaGroupId,
		"auto.offset.reset": "smallest"})
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
		return nil, err
	}
	return c, nil
}

func (l *Listener) SubscribeAndConsume() error {
	defer l.consumer.Close()
	err := l.consumer.SubscribeTopics(l.topics, nil)
	if err != nil {
		return err
	}

	// Graceful exit signalling
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true
	for run {
		select {
		case _ = <-sigchan:
			run = false
		default:
			msg, err := l.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			l.handleTaskByTopic(msg.TopicPartition.Topic, msg)
		}
	}
	log.Println("Kafka event subscriptions closed")
	return nil
}

// Routing should take a kafka message and enqueue onto processing queue
func (l *Listener) handleTaskByTopic(topic *string, msg *kafka.Message) {
	switch *topic {
	case "video-upload":
		log.Println("requesting enqueuing of inference on model task")
		l.processingQueue.Enqueue(tasks.InferenceOnModelTaskName, msg)
		return
	default:
		return
	}
}
