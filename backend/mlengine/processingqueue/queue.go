package processingqueue

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/hamdaankhalid/mlengine/dal"
	"github.com/hamdaankhalid/mlengine/database"
	"github.com/hamdaankhalid/mlengine/tasks"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type IQueue interface {
	BeginProcessing()
	Enqueue(taskName string, task interface{})
}

type Queue struct {
	uploadModelTasks      chan *tasks.UploadModelReq
	inferenceOnModelTasks chan *kafka.Message
}

func NewQueue() *Queue {
	return &Queue{
		uploadModelTasks:      make(chan *tasks.UploadModelReq),
		inferenceOnModelTasks: make(chan *kafka.Message)}
}

func (q *Queue) BeginProcessing() {
	// each task queue gets its own thread for execution
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		q.process(tasks.UploadModelTaskName)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		q.process(tasks.InferenceOnModelTaskName)
	}()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		defer wg.Done()
		running := true
		for running {
			select {
			case _ = <-sigchan:
				close(q.inferenceOnModelTasks)
				close(q.uploadModelTasks)
				running = false
			default:
				continue
			}
		}
	}()
	wg.Wait()
}

// Task should be a pointer passed val
func (q *Queue) Enqueue(taskName string, task interface{}) {
	switch taskName {
	case tasks.UploadModelTaskName:
		q.uploadModelTasks <- task.(*tasks.UploadModelReq)
		break
	case tasks.InferenceOnModelTaskName:
		q.inferenceOnModelTasks <- task.(*kafka.Message)
		break
	default:
		log.Println("No task queue corresponding to taskName: ", taskName)
		return
	}
	log.Printf("Enqueued Task: %s", taskName)
}

func (q *Queue) process(taskName string) {
	log.Printf("Task queue for task: %s starting", taskName)

	// Graceful exit signalling
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	conn, err := database.OpenConnection()
	if err != nil {
		log.Println("Task Not Queued Because No DB Connection:", taskName)
		return
	}

	run := true
	for run {
		select {
		case _ = <-sigchan:
			run = false
		default:
			switch taskName {
			case tasks.UploadModelTaskName:
				uploadModelReq := <-q.uploadModelTasks
				if uploadModelReq == nil {
					continue
				}
				queries := &dal.Queries{Conn: conn}
				tasks.UploadModelTask(uploadModelReq, queries)
			case tasks.InferenceOnModelTaskName:
				inferenceOnModelTask := <-q.inferenceOnModelTasks
				if inferenceOnModelTask == nil {
					continue
				}
				queries := &dal.Queries{Conn: conn}
				tasks.ParallelModelInferenceTask(inferenceOnModelTask, queries)
			default:
				log.Println("No corresponding task queue for task: ", taskName)
			}
		}
	}

	log.Printf("Task queue for task: %s shutdown", taskName)
}
