package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rs/zerolog/log"

	"kafka_consumer/pkg/api/bus"
	"kafka_consumer/pkg/service"
)

const (
	kafkaTopic = "test"
	kafkaDlx   = "test_dlx"
)

func main() {
	kcfg := &kafka.ConfigMap{
		"bootstrap.servers": "192.168.99.100:32768,192.168.99.100:32769,192.168.99.100:32770",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	}

	publisher, err := bus.NewPublisher(kcfg)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create publisher")
	}

	svc := service.New(publisher)

	consumer, err := bus.NewSubscriber(kcfg, svc)
	if err != nil {
		log.Panic().Err(err).Msg("Failed to create subscriber")
	}

	done := make(chan bool, 1)
	errs := make(chan error)
	go consumer.Start(kafkaTopic, errs)
	go publisher.Start(kafkaDlx, errs)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		<-c
		close(done)
	}()

	go func() {
		for {
			log.Err(<-errs).Msg("got new error")
		}
	}()

	<-done
	if err := consumer.Stop(); err != nil {
		log.Panic().Err(err).Msg("Failed to stop subscriber")
	}
	publisher.Stop()
	log.Error().Err(err).Msg("Kafka consumer terminated")
}
