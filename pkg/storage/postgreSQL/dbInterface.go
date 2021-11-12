package postgreSQL

import "CQRS-simple/pkg/models"

type DBInterface interface {
	CreateUser(u models.User) (models.User, error)
	Get(id string) (models.Read, error)
	CreatePost(p models.Post) (models.Post, error)
	CreateReadInfo(r models.User) error
	GetUser(id string) (models.User, error)
	AddPostToUserRead(r models.Read) error
	GetPosts(userID string) ([]models.PostRead, error)
}
