package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/hamdaankhalid/mlengine/database"
	"github.com/hamdaankhalid/mlengine/handlers"
	"github.com/hamdaankhalid/mlengine/kafkaworkers"
	"github.com/hamdaankhalid/mlengine/processingqueue"

	"github.com/joho/godotenv"
)

func setup() error {
	godotenv.Load(".env")
	log.Println("Initializing DB")
	err := database.InitDb(false)
	return err
}

func createAndStartServer(processingQueue *processingqueue.Queue) *http.Server {
	var wg *sync.WaitGroup
	router := handlers.NewRouter(processingQueue)
	server := &http.Server{Addr: ":6969", Handler: router.Routing}
	go func() {
		// let main know we are done cleaning up
		defer wg.Done()
		// always returns error. ErrServerClosed on graceful close
		// Blocking call so our above kicked off go routine will not exit early :)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			// log.Fatalln(err)
			return
		}
	}()
	log.Println("Listening And Serving ML Engine Up")
	return server
}

func start(wg *sync.WaitGroup) (*http.Server, error) {

	processingQueue := processingqueue.NewQueue()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		processingQueue.BeginProcessing()
		log.Println("Async processing shutdown on main thread")
	}(wg)

	// Start Event Consumer async
	eventListener, err := kafkaworkers.NewListener(processingQueue)
	if err != nil {
		return nil, err
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := eventListener.SubscribeAndConsume()
		if err != nil {
			log.Fatalf("Event Consumption Exited With Error: %s", err)
		}
		log.Println("Kafka event processing shutdown on main thread")
	}(wg)

	//Start Http server async
	server := createAndStartServer(processingQueue)
	return server, nil
}

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Use waitgroups for graceful exit on http server
	var wg sync.WaitGroup

	_, err = start(&wg)
	wg.Wait()

	if err != nil {
		log.Fatal(err)
	}
}
