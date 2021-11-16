package inMemory

import "CQRS-simple/pkg/models"

type InMemoryInterface interface {
	CreateUser(ur models.Read) error
	CreatePost(ur models.Read) error
	GetAllUsers() (*[]models.User, error)
	GetUserPosts(id string) (*models.UserPosts, error)
}
