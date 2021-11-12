package command

import "CQRS-simple/pkg/models"

type CommandInterface interface {
	CreateUser(u models.User) (*models.User, error)
	CreatePost(p models.Post) (*models.Post, error)
}
