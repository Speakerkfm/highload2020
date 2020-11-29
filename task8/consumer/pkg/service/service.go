package service

import (
	"github.com/rs/zerolog/log"

	"kafka_consumer/pkg/models"
)

type publisher interface {
	Publish(msg models.Message)
}

type service struct {
	pub publisher
}

func New(pub publisher) models.Service {
	return &service{
		pub: pub,
	}
}

func (s *service) HandleMessage(msg models.Message) error {
	if msg.Type == "message" {
		log.Info().Msg("Done")
		return nil
	}

	s.pub.Publish(msg)

	return nil
}
