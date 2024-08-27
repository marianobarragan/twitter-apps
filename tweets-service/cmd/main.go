package main

import (
	"fmt"
	"log"
	"time"
	"tweets-service/domain"
	"tweets-service/events"
	"tweets-service/routes"
	"tweets-service/storage/gomemdb"
)

func main() {
	time.Sleep(15 * time.Second)
	repository, err := gomemdb.NewRepository()
	if err != nil {
		log.Fatalf("Error running tweets service - err %s \n", err)
		return
	}
	eventProducer, err := events.NewEventProducer()
	if err != nil {
		log.Fatalf("Error running tweets service - err %s \n", err)
		return
	}
	defer eventProducer.Close()
	s := domain.NewService(repository, nil)
	router := routes.NewRouter(s)
	err = router.Run(":8082")
	if err != nil {
		log.Fatalf("Error running tweets service - err %s \n", err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	fmt.Println("Hello, World! from tweets service")
}
