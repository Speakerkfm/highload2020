package models

type Service interface {
	HandleMessage(msg Message) error
}
