package readServ

import "CQRS-simple/pkg/models"

type ReadServInterface interface {
	GetAllUsers() (*[]models.User, error)
	UserPosts(userID string) (models.UserPosts, error)
}
