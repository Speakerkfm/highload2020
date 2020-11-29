package bus

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"

	"kafka_consumer/pkg/models"
)

const saveEntity string = "save_entity"

type Subscriber struct {
	consumer *kafka.Consumer
	callback func(msg models.Message) error
}

func NewSubscriber(configMap *kafka.ConfigMap, svc models.Service) (*Subscriber, error) {
	c, err := kafka.NewConsumer(configMap)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	return &Subscriber{
		consumer: c,
		callback: svc.HandleMessage,
	}, nil
}

func (s *Subscriber) Start(topic string, errs chan<- error) {
	if err := s.consumer.Subscribe(topic, nil); err != nil {
		errs <- fmt.Errorf("consumer error: %v\n", err)
		return
	}

	for {
		m, err := s.consumer.ReadMessage(-1)
		if err != nil {
			errs <- fmt.Errorf("consumer error: %v (%v)\n", err, m)
			continue
		}

		var msg models.Message
		if err := json.Unmarshal(m.Value, &msg); err != nil {
			errs <- fmt.Errorf("consumer error: %v (%v)\n", err, msg)
			continue
		}
		if err := s.callback(msg); err != nil {
			log.Error().Err(err).Msgf("consumer error wile processing message %v", msg)
		}
	}
}

func (s *Subscriber) Stop() error {
	return s.consumer.Close()
}
