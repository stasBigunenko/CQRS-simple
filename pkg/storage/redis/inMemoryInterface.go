package redis

import "CQRS-simple/pkg/models"

type RedisDBInterface interface {
	CreateUser(ur models.Read) (models.User, error)
	CreatePost(ur models.Post) (models.PostRead, error)
	GetAllUsers() (*[]models.User, error)
	GetUserPosts(id string) (models.UserPosts, error)
	UpdateUser(u models.User) error
	UpdatePost(p models.Post) error
	DeleteUser(id string) error
	DeletePost(id, userID string) error
	Exist(id string) bool
}
