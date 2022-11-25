package workers

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

/**
	Listener is created via NewListener, it runs a blocking method
	which listens to incoming topic messages and processes them in
	it's own thread.
**/
type Listener struct {
	consumer *kafka.Consumer
	topics   []string
}

func NewListener() (*Listener, error) {
	topics := []string{"video-upload"}

	consumer, err := initConsumer()
	if err != nil {
		return nil, err
	}
	return &Listener{consumer: consumer, topics: topics}, nil
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
			// Terminate
			run = false
			os.Exit(0)
		default:
			msg, err := l.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			handler := RouteTask(msg.TopicPartition.Topic)
			go handler(msg)
		}
	}
	return nil
}
