package main

import (
	"log"
	"timeline-service/domain"
	"timeline-service/routes"
	"timeline-service/storage/gomemdb"
)

func main() {
	repository, err := gomemdb.NewRepository()
	if err != nil {
		log.Fatalf("Error running timeline service storage - err %s \n", err)
		return
	}
	s := domain.NewService(repository)
	router := routes.NewRouter(s)
	err = router.Run(":8081")
	if err != nil {
		log.Fatalf("Error running timeline service - err %s \n", err)
		return
	}
	log.Println("Users service is running!")
}
