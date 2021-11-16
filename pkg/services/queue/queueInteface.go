package queue

import "CQRS-simple/pkg/models"

type QueueInterface interface {

	//UserPosts(userID string) (*models.UserPosts, error)
	//GetAllUsers() (*[]models.User, error)

	//Get(id string) (*models.Read, error)
	//GetUser(id string) (*models.User, error)
	//GetPost(id string) (*models.Post, error)

	GetAllUsers() (*[]models.User, error)
	UserPosts(userID string) (*models.UserPosts, error)
}
