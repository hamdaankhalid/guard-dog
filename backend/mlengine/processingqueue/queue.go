package processingqueue

import (
	"log"
	"mime/multipart"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

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

type Queue struct {
	uploadModelTasks      chan *UploadModelReq
	inferenceOnModelTasks chan *kafka.Message
}

func NewQueue() *Queue {
	return &Queue{
		uploadModelTasks:      make(chan *UploadModelReq),
		inferenceOnModelTasks: make(chan *kafka.Message)}
}

func (q *Queue) BeginProcessing() {
	// each task queue gets its own thread for execution
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		q.process(UploadModelTaskName)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		q.process(InferenceOnModelTaskName)
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
	case UploadModelTaskName:
		q.uploadModelTasks <- task.(*UploadModelReq)
		break
	case InferenceOnModelTaskName:
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

	run := true
	for run {
		select {
		case _ = <-sigchan:
			log.Println("Exiting async processing loop")
			run = false
		default:
			switch taskName {
			case UploadModelTaskName:
				uploadModelReq := <-q.uploadModelTasks
				if uploadModelReq == nil {
					continue
				}
				uploadModelTask(uploadModelReq)
			case InferenceOnModelTaskName:
				inferenceOnModelTask := <-q.inferenceOnModelTasks
				if inferenceOnModelTask == nil {
					continue
				}
				parallelModelInferenceTask(inferenceOnModelTask)
			default:
				log.Println("No corresponding task queue for task: ", taskName)
			}
		}
	}

	log.Printf("Task queue for task: %s shutdown", taskName)
}
