package queue

import "CQRS-simple/pkg/models"

type QueueInterface interface {
	//Get(id string) (*models.Read, error)
	UserPosts(userID string) (*models.UserPosts, error)
	GetAllUsers() (*[]models.User, error)
	GetUser(id string) (*models.Read, error)
	GetPost(id string) (*models.Read, error)
}
