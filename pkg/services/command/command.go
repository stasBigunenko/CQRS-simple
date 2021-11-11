package command

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/postgreSQL"
	"errors"
)

type Command struct {
	command postgreSQL.DBInterface
}

func NewCommand(c postgreSQL.DBInterface) Command {
	return Command{
		command: c,
	}
}

func (s *Command) Create(u models.User) (*models.User, error) {

	if u.Name == "" || u.Surname == "" || u.Age == 0 || u.Sex == "" {
		return nil, errors.New("some fields are missing")
	}

	userNew, err := s.command.Create(u)
	if err != nil {
		return nil, err
	}

	return &userNew, nil
}
