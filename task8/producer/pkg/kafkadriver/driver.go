package kafkadriver

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const flushTimeout = 15 * 100

type KafkaMsg struct {
	Type string `json:"type"`
}

func CreateProducer(topic string) (chan<- KafkaMsg, <-chan string, <-chan error) {
	errs := make(chan error, 1)
	res := make(chan KafkaMsg)
	notifications := make(chan string)

	go func() {
		defer close(res)

		p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.99.100:32768,192.168.99.100:32769,192.168.99.100:32770"})
		if err != nil {
			errs <- fmt.Errorf("failed to create producer: %w", err)
			return
		}
		defer func() {
			p.Flush(flushTimeout)
			p.Close()
		}()

		// Delivery report handler for produced messages
		go func() {
			for e := range p.Events() {
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

		for msg := range res {
			value, _ := json.Marshal(msg)
			if err := p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          value,
			}, nil); err != nil {
				errs <- err
				return
			}
		}
	}()

	return res, notifications, errs
}