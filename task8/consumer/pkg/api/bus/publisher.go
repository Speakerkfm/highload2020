package bus

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"

	"kafka_consumer/pkg/models"
)

const flushTimeout = 15 * 100

type publisher struct {
	producer *kafka.Producer
	ch       chan models.Message
}

func NewPublisher(configMap *kafka.ConfigMap) (*publisher, error) {
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		return nil, errors.Wrap(err, "fail to craete producer")
	}
	return &publisher{
		producer: p,
		ch:       make(chan models.Message),
	}, nil
}

func (p *publisher) Start(topic string, errs chan<- error) {
	notifications := make(chan string)

	// Delivery report handler for produced messages
	go func() {
		for e := range p.producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					notifications <- fmt.Sprintf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					notifications <- fmt.Sprintf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	go func() {
		for msg := range p.ch {
			value, _ := json.Marshal(msg)
			if err := p.producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          value,
			}, nil); err != nil {
				errs <- err
				return
			}
		}
	}()
}

func (p *publisher) Publish(msg models.Message) {
	p.ch <- msg
}

func (p *publisher) Stop() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
	close(p.ch)
}
