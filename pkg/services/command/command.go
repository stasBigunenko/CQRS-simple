package command

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/inMemory"
	"CQRS-simple/pkg/storage/postgreSQL"
	"errors"
)

type Command struct {
	command postgreSQL.DBInterface
	storage inMemory.InMemoryInterface
}

func NewCommand(c postgreSQL.DBInterface, s inMemory.InMemoryInterface) Command {
	return Command{
		command: c,
		storage: s,
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

	//c.command.CreateReadInfo(r)
	c.storage.CreateUser(r)

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

	//userRead, err := c.command.GetUserRead(p.UserID)
	userRead, err := c.command.GetUser(p.UserID)

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

	//c.command.CreateReadInfo(r)
	c.storage.CreatePost(r)

	return &postNew, nil
}

func (c *Command) UpdateUser(u models.User) (*models.User, error) {

	userNew, err := c.command.UpdateUser(u)
	if err != nil {
		return &models.User{}, err
	}

	//userRead, err := c.command.GetUserRead(u.ID)
	//
	//userRead.User.ID = u.ID
	//userRead.User.Name = userNew.Name
	//userRead.User.Age = userNew.Age

	//c.command.CreateReadInfo(userRead)
	c.command.UpdateReadUser(userNew)

	return &userNew, nil
}

func (c *Command) UpdatePost(p models.Post) (*models.Post, error) {

	postNew, err := c.command.UpdatePost(p)
	if err != nil {
		return nil, err
	}

	//userRead, err := c.command.GetUserRead(p.UserID)
	//
	//user := models.User{
	//	ID:   userRead.User.ID,
	//	Name: userRead.User.Name,
	//	Age:  userRead.User.Age,
	//}
	//
	//postRead := models.PostRead{
	//	ID:      postNew.ID,
	//	Title:   postNew.Title,
	//	Message: postNew.Message,
	//}
	//
	//r := models.Read{
	//	User:     user,
	//	PostRead: postRead,
	//}

	//c.command.CreateReadInfo(r)
	c.command.UpdateReadPost(postNew)

	return &postNew, nil
}

func (c *Command) DeleteUser(id string) error {
	err := c.command.DeleteUser(id)
	if err != nil {
		return err
	}

	err = c.command.DeleteReadUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (c *Command) DeletePost(id string) error {
	err := c.command.DeletePost(id)
	if err != nil {
		return err
	}

	err = c.command.DeleteReadPost(id)
	if err != nil {
		return err
	}

	return nil
}
