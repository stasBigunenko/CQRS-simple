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

func (q *Queue) Get(id string) (*models.User, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("service: couldn't parse id")
	}

	user, err := q.queue.Get(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
