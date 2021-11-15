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

func (c *Command) CreateUser(u models.User) (*models.User, error) {

	if u.Name == "" || u.Age == 0 {
		return nil, errors.New("some fields are missing")
	}

	userNew, err := c.command.CreateUser(u)
	if err != nil {
		return nil, err
	}

	var r models.Read

	r.User.ID = userNew.ID
	r.User.Name = userNew.Name
	r.User.Age = userNew.Age
	r.PostRead.ID = "empty"
	r.PostRead.Title = "empty"
	r.PostRead.Message = "empty"

	c.command.CreateReadInfo(r)

	return &userNew, nil
}

func (c *Command) CreatePost(p models.Post) (*models.Post, error) {

	if p.Title == "" || p.Message == "" {
		return nil, errors.New("some fields are missing")
	}

	postNew, err := c.command.CreatePost(p)
	if err != nil {
		return nil, err
	}

	userRead, err := c.command.GetUserRead(p.UserID)

	user := models.User{
		ID:   userRead.User.ID,
		Name: userRead.User.Name,
		Age:  userRead.User.Age,
	}

	postRead := models.PostRead{
		ID:      postNew.ID,
		Title:   postNew.Title,
		Message: postNew.Message,
	}

	r := models.Read{
		User:     user,
		PostRead: postRead,
	}

	c.command.CreateReadInfo(r)

	return &postNew, nil
}
