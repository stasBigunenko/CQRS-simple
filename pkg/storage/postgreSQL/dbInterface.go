package postgreSQL

import "CQRS-simple/pkg/models"

type DBInterface interface {
	Create(u models.User) (models.User, error)
	Get(id string) (models.User, error)
}
