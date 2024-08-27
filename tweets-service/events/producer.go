package events

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"strconv"
	"time"
	"tweets-service/domain"
)

const (
	KafkaServerAddress = "localhost:9092"
	KafkaTopic         = "notifications"
)

type RealEventProducer struct {
	producer sarama.SyncProducer
}

var _ domain.EventProducer = RealEventProducer{}

func NewEventProducer() (domain.EventProducer, error) {
	producer, err := setupProducer()
	if err != nil {
		return RealEventProducer{}, err
	}
	return RealEventProducer{producer: producer}, nil
}

func setupProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{KafkaServerAddress}, config)
	if err != nil {
		return nil, fmt.Errorf("failed to setup producer: %w", err)
	}
	return producer, nil
}

func (r RealEventProducer) PublishNewTweetEvent(userID int, tweet domain.Tweet) error {
	event := NewTweetEvent{
		TweetID:  tweet.ID,
		AuthorID: userID,
	}
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: KafkaTopic,
		Key:   sarama.StringEncoder(strconv.Itoa(int(time.Now().UnixMilli()))),
		Value: sarama.StringEncoder(eventJSON),
	}
	_, _, err = r.producer.SendMessage(msg)
	return err
}

func (r RealEventProducer) Close() error {
	return r.producer.Close()
}
