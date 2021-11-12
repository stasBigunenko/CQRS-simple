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

	c.command.CreateReadInfo(userNew)

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

	user, err := c.command.GetUser(p.UserID)

	postRead := models.PostRead{
		ID:      postNew.ID,
		Title:   postNew.Title,
		Message: postNew.Message,
	}

	r := models.Read{
		User:     user,
		PostRead: postRead,
	}

	c.command.AddPostToUserRead(r)

	return &postNew, nil
}

//func (c *Command ) UserPosts (userID string) (*models.UserPosts, error) {
//
//	user, err := c.command.GetUser(userID)
//	if err != nil {
//		return &models.UserPosts{}, err
//	}
//
//	postsRead, err := c.command.GetPosts(userID)
//	if err != nil {
//		return &models.UserPosts{}, err
//	}
//
//	userPosts := models.UserPosts{
//		user,
//		postsRead,
//	}
//
//	return &userPosts, nil
//}
