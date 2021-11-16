package postgreSQL

import "CQRS-simple/pkg/models"

type DBInterface interface {
	// Write
	CreateUser(u models.User) (models.User, error) // Table 1
	CreatePost(p models.Post) (models.Post, error) // Table 2
	UpdateUser(u models.User) (models.User, error)
	UpdatePost(p models.Post) (models.Post, error)
	DeleteUser(id string) error
	DeletePost(id string) error

	// Read
	//Get(id string) (models.Read, error)
	GetUserRead(id string) (models.Read, error)
	GetPostRead(id string) (models.Read, error)
	GetAllUsers() (*[]models.User, error)
	GetPosts(userID string) ([]models.PostRead, error)

	// Help
	GetUser(id string) (models.Read, error)
	CreateReadInfo(r models.Read) error
	DeleteReadUser(id string) error
	DeleteReadPost(id string) error
	UpdateReadUser(u models.User) error
	UpdateReadPost(p models.Post) error
}
