package events

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"timeline-subscriber/domain"
)

func NewSubscriber(service domain.Service) Subscriber {
	eventsChan := make(chan NewTweetEvent)
	ctx, cancel := context.WithCancel(context.Background())
	go setupConsumerGroup(ctx, eventsChan)
	defer cancel()
	return RealSubscriber{
		service:    service,
		eventsChan: eventsChan,
	}
}

type Subscriber interface {
	ConsumeEvents() error
}

type RealSubscriber struct {
	service    domain.Service
	consumer   consumer
	eventsChan chan NewTweetEvent
}

func (subscriber RealSubscriber) ConsumeEvents() error {
	var (
		event NewTweetEvent
		err   error
	)
	for {
		event, err = subscriber.GetNewEvent()
		if err != nil {
			break
		}
		err = subscriber.service.IndexTweetTimeline(event.AuthorID)
		if err != nil {
			break
		}
	}
	return err
}

func (subscriber RealSubscriber) GetNewEvent() (event NewTweetEvent, err error) {
	event = <-subscriber.eventsChan
	return event, nil
}

const (
	ConsumerGroup      = "notifications-group"
	ConsumerTopic      = "notifications"
	KafkaServerAddress = "localhost:9092"
)

// ============== KAFKA RELATED FUNCTIONS ==============
type consumer struct {
	eventsChan chan NewTweetEvent
}

func (*consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)
		log.Printf("getting notification - id: %s", userID)
		var notification NewTweetEvent
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}
		consumer.eventsChan <- notification
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{KafkaServerAddress}, ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func setupConsumerGroup(ctx context.Context, eventsChan chan NewTweetEvent) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
		panic(err)
	}
	defer consumerGroup.Close()

	c := &consumer{
		eventsChan: eventsChan,
	}
	fmt.Printf("Kafka CONSUMER (Group: %s) 👥📥 started at http://localhost%s\n", ConsumerGroup)
	for {
		err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, c)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}
