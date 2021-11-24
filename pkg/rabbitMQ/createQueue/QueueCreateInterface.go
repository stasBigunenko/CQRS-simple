package createQueue

import "CQRS-simple/pkg/models"

type QueueCreateInterface interface {
	QueueCreateWrite(c models.Cud) error
	QueueCreateCache(up models.Cud) error
}
