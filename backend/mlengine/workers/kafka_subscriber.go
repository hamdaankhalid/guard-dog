package workers

import (
	"fmt"
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
	// Create Consumer instance
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

	err := l.consumer.SubscribeTopics(l.topics, nil)
	defer l.consumer.Close()

	if err != nil {
		return err
	}
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	run := true
	fmt.Println("Listening")
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			msg, err := l.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			recordKey := string(msg.Key)
			fmt.Println(recordKey)
			handler := RouteTask(recordKey)
			go handler(msg)
		}
	}

	return nil
}
