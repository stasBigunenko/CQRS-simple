package inMemory

import "CQRS-simple/pkg/models"

type InMemoryInterface interface {
	CreateUser(ur models.Read) error
	CreatePost(ur models.Read) error
	GetAllUsers() (*[]models.User, error)
	GetUserPosts(id string) (*models.UserPosts, error)
	UpdateUser(u models.User) error
	UpdatePost(p models.Post) error
	DeleteUser(id string) error
	DeletePost(id, userID string) error
}
