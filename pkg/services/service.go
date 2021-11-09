package services

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/postgreSQL"
	"errors"
	"github.com/google/uuid"
)

type Services struct {
	service postgreSQL.DBInterface
}

func NewService(s postgreSQL.DBInterface) Services {
	return Services{
		service: s,
	}
}

func (s *Services) Create(u models.User) (*models.User, error) {

	if u.Name == "" || u.Surname == "" || u.Age == 0 || u.Sex == "" {
		return nil, errors.New("some fields are missing")
	}

	userNew, err := s.service.Create(u)
	if err != nil {
		return nil, err
	}
	return &userNew, nil
}

func (s *Services) Get(id string) (*models.User, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("service: couldn't parse id")
	}

	user, err := s.service.Get(id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
