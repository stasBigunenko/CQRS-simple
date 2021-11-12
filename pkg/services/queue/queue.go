package queue

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/postgreSQL"
	"errors"
	"github.com/google/uuid"
)

type Queue struct {
	queue postgreSQL.DBInterface
}

func NewQueue(q postgreSQL.DBInterface) Queue {
	return Queue{
		queue: q,
	}
}

func (q *Queue) Get(id string) (*models.Read, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("service: couldn't parse id")
	}

	read, err := q.queue.Get(id)
	if err != nil {
		return nil, err
	}

	return &read, nil
}

func (q *Queue) UserPosts(userID string) (*models.UserPosts, error) {

	user, err := q.queue.GetUser(userID)
	if err != nil {
		return &models.UserPosts{}, err
	}

	postsRead, err := q.queue.GetPosts(userID)
	if err != nil {
		return &models.UserPosts{}, err
	}

	userPosts := models.UserPosts{
		user,
		postsRead,
	}

	return &userPosts, nil
}
