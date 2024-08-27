package main

import (
	"log"
	"users-service/domain"
	"users-service/routes"
	"users-service/storage/gomemdb"
)

func main() {
	repository, err := gomemdb.NewRepository()
	if err != nil {
		log.Fatalf("Error running users service - err %s \n", err)
		return
	}
	s := domain.NewService(repository)
	router := routes.NewRouter(s)
	err = router.Run(":8083")
	if err != nil {
		log.Fatalf("Error running users service - err %s \n", err)
		return
	}
	log.Println("Users service is running!")
}
