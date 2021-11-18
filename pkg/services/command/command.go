package command

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/createQueue"
	"CQRS-simple/pkg/storage/inMemory"
	"CQRS-simple/pkg/storage/postgreSQL"
	"errors"
	"github.com/google/uuid"
	"log"
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

	//c.command.CreateReadInfo(r) function for inMemory Storage
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

	//userRead, err := c.command.GetUserRead(p.UserID) function for inMemory Storage
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

	//c.command.CreateReadInfo(r) function for inMemory Storage
	c.storage.CreatePost(r)

	return &postNew, nil
}

func (c *Command) UpdateUser(u models.User) (*models.User, error) {

	userNew, err := c.command.UpdateUser(u)
	if err != nil {
		return &models.User{}, err
	}

	//c.command.UpdateReadUser(userNew) function for inMemory Storage
	exist := c.storage.Exist(u.ID)
	if exist {
		var cud models.Cud
		cud.Model = "user"
		cud.Command = "update"
		cud.User = userNew
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE UPDATE USER SENDED-------------")
		//err = c.storage.UpdateUser(userNew)
		//if err != nil {
		//	return &models.User{}, err
		//}
	}

	return &userNew, nil
}

func (c *Command) UpdatePost(p models.Post) (*models.Post, error) {

	postNew, err := c.command.UpdatePost(p)
	if err != nil {
		return nil, err
	}

	//c.command.UpdateReadPost(postNew) function for inMemory Storage
	exist := c.storage.Exist(postNew.UserID)
	if exist {
		var cud models.Cud
		cud.Model = "post"
		cud.Command = "update"
		cud.Post = postNew
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE UPDATE POST SENDED-------------")
		//err = c.storage.UpdatePost(postNew)
		//if err != nil {
		//	return &models.Post{}, err
		//}
	}

	return &postNew, nil
}

func (c *Command) DeleteUser(id string) error {
	err := c.command.DeleteUser(id)
	if err != nil {
		return err
	}

	//err = c.command.DeleteReadUser(id) function for inMemory Storage
	//if err != nil {
	//	return err
	//}
	exist := c.storage.Exist(id)
	if exist {
		var cud models.Cud
		cud.Model = "user"
		cud.Command = "delete"
		cud.User.ID = id
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE DELETE USER SENDED-------------")
		//err = c.storage.DeleteUser(id)
		//if err != nil {
		//	return err
		//}
	}
	return nil
}

func (c *Command) DeletePost(id string) error {

	mp, err := c.GetPost(id)
	if err != nil {
		return err
	}

	err = c.command.DeletePost(id)
	if err != nil {
		return err
	}

	//err = c.command.DeleteReadPost(id) function for inMemory Storage
	//if err != nil {
	//	return err
	//}

	userID := mp.UserID

	exist := c.storage.Exist(userID)
	if exist {
		var cud models.Cud
		cud.Model = "post"
		cud.Command = "delete"
		cud.Post = *mp
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE DELETE POST SENDED-------------")
		//err = c.storage.DeletePost(id, userID)
		//if err != nil {
		//	return err
		//}
	}

	return nil
}

func (c *Command) GetPost(id string) (*models.Post, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("service: couldn't parse id")
	}

	read, err := c.command.GetPostRead(id)
	if err != nil {
		return nil, err
	}

	post := models.Post{
		ID:      read.PostRead.ID,
		UserID:  read.User.ID,
		Title:   read.PostRead.Title,
		Message: read.PostRead.Message,
	}

	return &post, nil
}
