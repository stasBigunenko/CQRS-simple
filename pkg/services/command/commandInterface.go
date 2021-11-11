package command

import "CQRS-simple/pkg/models"

type CommandInterface interface {
	Create(u models.User) (*models.User, error)
}
