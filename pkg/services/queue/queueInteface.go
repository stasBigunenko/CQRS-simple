package queue

import "CQRS-simple/pkg/models"

type QueueInterface interface {
	Get(id string) (*models.Read, error)
	UserPosts(userID string) (*models.UserPosts, error)
}
