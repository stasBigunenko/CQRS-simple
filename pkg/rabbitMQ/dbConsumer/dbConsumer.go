package dbConsumer

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/services/readServ"
	"CQRS-simple/pkg/services/writeServ"
	"log"
)

type DBConsumer struct {
	command writeServ.WriteServInterface
	queue   readServ.ReadServInterface
}

func NewDBConsumer(c writeServ.WriteServInterface, q readServ.ReadServInterface) *DBConsumer {
	return &DBConsumer{
		command: c,
		queue:   q,
	}
}

func (n *DBConsumer) Received(cud models.Cud) {
	switch cud.Model {
	case "user":
		n.ReceivedUser(cud)
	case "post":
		n.ReceivedPost(cud)
	}
}

func (n *DBConsumer) ReceivedUser(cud models.Cud) {
	switch cud.Command {
	case "create":
		n.CreateUser(cud)
	case "update":
		n.UpdateUser(cud)
	case "delete":
		n.DeleteUser(cud)
	}
}

func (n *DBConsumer) ReceivedPost(cud models.Cud) {
	switch cud.Command {
	case "create":
		n.CreatePost(cud)
	case "update":
		n.UpdatePost(cud)
	case "delete":
		n.DeletePost(cud)
	}
}

func (n *DBConsumer) CreateUser(cud models.Cud) {
	var user models.User

	user = cud.User
	_, err := n.command.CreateUser(user)
	if err != nil {
		log.Printf("queu cud problems %s\n", err)
		return
	}
}

func (n *DBConsumer) CreatePost(cud models.Cud) {
	var post models.Post

	post = cud.Post
	_, err := n.command.CreatePost(post)
	if err != nil {
		log.Printf("queu cud problems %s\n", err)
		return
	}
}

func (n *DBConsumer) UpdateUser(cud models.Cud) {
	var user models.User

	user = cud.User
	_, err := n.command.UpdateUser(user)
	if err != nil {
		log.Printf("queu cud problems %s\n", err)
		return
	}
}

func (n *DBConsumer) UpdatePost(cud models.Cud) {
	var post models.Post

	post = cud.Post
	_, err := n.command.UpdatePost(post)
	if err != nil {
		log.Printf("queu cud problems %s\n", err)
		return
	}
}

func (n *DBConsumer) DeleteUser(cud models.Cud) {
	var user models.User

	user = cud.User
	err := n.command.DeleteUser(user.ID)
	if err != nil {
		log.Printf("queu cud problems %s\n", err)
		return
	}
}

func (n *DBConsumer) DeletePost(cud models.Cud) {
	var post models.Post

	post = cud.Post
	err := n.command.DeletePost(post.ID)
	if err != nil {
		log.Printf("queu cud problems %s\n", err)
		return
	}
}
