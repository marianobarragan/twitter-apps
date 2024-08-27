package main

import (
	"log"
	"time"
	"timeline-subscriber/domain"
	"timeline-subscriber/events"
	"timeline-subscriber/rest"
)

func main() {
	time.Sleep(15 * time.Second)
	tweetsClient := rest.NewTweetsClient()
	usersClient := rest.NewUsersClient()
	timelineClient := rest.NewTimelineClient()
	service := domain.NewService(timelineClient, tweetsClient, usersClient)
	subscriber := events.NewSubscriber(service)
	log.Println("Subscriber service is running!")
	err := subscriber.ConsumeEvents()
	if err != nil {
		log.Fatalf("Error running subscriber service - err %s \n", err)
		return
	}
}
