package command

import "CQRS-simple/pkg/models"

type CommandInterface interface {
	CreateUser(u models.User) (*models.User, error)
	CreatePost(p models.Post) (*models.Post, error)
	UpdateUser(u models.User) (*models.User, error)
	UpdatePost(p models.Post) (*models.Post, error)
	DeleteUser(id string) error
	DeletePost(id string) error
}
