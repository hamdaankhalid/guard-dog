package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

func start(wg *sync.WaitGroup) (*http.Server, error) {
	// Start Processing Queue asyn
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

	// Start Http server async
	router := handlers.NewRouter(processingQueue)
	server := &http.Server{Addr: ":6969", Handler: router.Routing}
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln(err)
			return
		}
		log.Println("Http Web-Server Shutdown")
	}()

	log.Println("Listening And Serving ML Engine Up")

	return server, nil
}

func main() {
	err := setup()
	if err != nil {
		log.Fatal(err)
		return
	}
	var wg sync.WaitGroup
	server, err := start(&wg)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go func(server *http.Server, wg *sync.WaitGroup) {
		defer wg.Done()
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		for true {
			select {
			case _ = <-sigchan:
				server.Shutdown(nil)
				return
			}
		}
	}(server, &wg)
	wg.Wait()
}
