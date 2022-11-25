package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/hamdaankhalid/mlengine/database"
	"github.com/hamdaankhalid/mlengine/workers"

	"github.com/joho/godotenv"
)

func setup() error {
	godotenv.Load(".env")
	log.Println("Initializing DB")
	err := database.InitDb(false)
	return err
}

func createAndStartServer(wg *sync.WaitGroup) *http.Server {
	server := &http.Server{Addr: ":6969", Handler: GetRouter()}
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
	// Start Event Consumer async
	listener, err := workers.NewListener()
	if err != nil {
		return nil, err
	}
	go func() {
		err := listener.SubscribeAndConsume()
		if err != nil {
			log.Fatalf("Event Consumption Exited With Error: %s", err)
		}
	}()

	//Start Http server async
	server := createAndStartServer(wg)

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
	wg.Add(1)
	_, err = start(&wg)
	wg.Wait()

	if err != nil {
		log.Fatal(err)
	}
}
