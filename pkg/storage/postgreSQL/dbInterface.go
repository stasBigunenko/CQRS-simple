package postgreSQL

import "CQRS-simple/pkg/models"

type DBInterface interface {
	// Write
	CreateUser(u models.User) (models.User, error) // Table 1
	CreatePost(p models.Post) (models.Post, error) // Table 2

	// Read
	//Get(id string) (models.Read, error)
	GetUser(id string) (models.User, error)
	GetPost(id string) (models.Post, error)
	GetAllUsers() (*[]models.User, error)
	GetPosts(userID string) ([]models.PostRead, error)

	// Help
	AddPostToUserRead(r models.Read) error
	CreateReadInfo(r models.User) error
}
