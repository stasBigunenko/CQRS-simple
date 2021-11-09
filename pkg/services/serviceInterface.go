package services

import "CQRS-simple/pkg/models"

type ServiceInterface interface {
	Create(u models.User) (*models.User, error)
	Get(id string) (*models.User, error)
}
