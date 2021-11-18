package queue

import "CQRS-simple/pkg/models"

type QueueInterface interface {
	GetAllUsers() (*[]models.User, error)
	UserPosts(userID string) (models.UserPosts, error)

	//UserPosts(userID string) (*models.UserPosts, error)
	//GetAllUsers() (*[]models.User, error)

	//GetUser(id string) (*models.User, error)
	//GetPost(id string) (*models.Post, error)
}
