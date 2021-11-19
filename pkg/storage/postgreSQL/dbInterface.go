package postgreSQL

import "CQRS-simple/pkg/models"

type DBInterface interface {
	// Write
	CreateUser(u models.User) (models.User, error) // Table 1
	CreatePost(p models.Post) (models.Post, error) // Table 2
	UpdateUser(u models.User) (models.User, error) // Table 1
	UpdatePost(p models.Post) (models.Post, error) // Table 2
	DeleteUser(id string) error                    // Table 1
	DeletePost(id string) error                    // Table 2

	// Read
	GetPostRead(id string) (models.Read, error)
	GetAllUsers() (*[]models.User, error)
	GetPosts(userID string) ([]models.PostRead, error)

	// Help
	GetUser(id string) (models.Read, error)
}
