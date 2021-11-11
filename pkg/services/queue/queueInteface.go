package queue

import "CQRS-simple/pkg/models"

type QueueInterface interface {
	Get(id string) (*models.User, error)
}
