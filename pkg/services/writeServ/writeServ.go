package writeServ

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/createQueue"
	"CQRS-simple/pkg/storage/postgreSQL"
	"CQRS-simple/pkg/storage/redis"
	"errors"
	"github.com/google/uuid"
	"log"
)

type WriteServ struct {
	postgresDB postgreSQL.DBInterface
	redisDB    redis.RedisDBInterface
}

func NewWriteServ(p postgreSQL.DBInterface, r redis.RedisDBInterface) WriteServ {
	return WriteServ{
		postgresDB: p,
		redisDB:    r,
	}
}

func (w *WriteServ) CreateUser(u models.User) (*models.User, error) {

	userNew, err := w.postgresDB.CreateUser(u)
	if err != nil {
		return nil, err
	}

	return &userNew, nil
}

func (w *WriteServ) CreatePost(p models.Post) (*models.Post, error) {

	postNew, err := w.postgresDB.CreatePost(p)
	if err != nil {
		return nil, err
	}

	exist := w.redisDB.Exist(postNew.UserID)
	if exist {
		var cud models.Cud
		cud.Model = "post"
		cud.Command = "create"
		cud.Post = postNew
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE CREATE POST SENDED-------------")
	}

	return &postNew, nil
}

func (w *WriteServ) UpdateUser(u models.User) (*models.User, error) {

	userNew, err := w.postgresDB.UpdateUser(u)
	if err != nil {
		return &models.User{}, err
	}

	exist := w.redisDB.Exist(u.ID)
	if exist {
		var cud models.Cud
		cud.Model = "user"
		cud.Command = "update"
		cud.User = userNew
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE UPDATE USER SENDED-------------")
	}

	return &userNew, nil
}

func (w *WriteServ) UpdatePost(p models.Post) (*models.Post, error) {

	postNew, err := w.postgresDB.UpdatePost(p)
	if err != nil {
		return nil, err
	}

	exist := w.redisDB.Exist(postNew.UserID)
	if exist {
		var cud models.Cud
		cud.Model = "post"
		cud.Command = "update"
		cud.Post = postNew
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE UPDATE POST SENDED-------------")
	}

	return &postNew, nil
}

func (w *WriteServ) DeleteUser(id string) error {
	err := w.postgresDB.DeleteUser(id)
	if err != nil {
		return err
	}

	exist := w.redisDB.Exist(id)
	if exist {
		var cud models.Cud
		cud.Model = "user"
		cud.Command = "delete"
		cud.User.ID = id
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE DELETE USER SENDED-------------")
	}
	return nil
}

func (w *WriteServ) DeletePost(id string) error {

	mp, err := w.GetPost(id)
	if err != nil {
		return err
	}

	err = w.postgresDB.DeletePost(id)
	if err != nil {
		return err
	}

	userID := mp.UserID

	exist := w.redisDB.Exist(userID)
	if exist {
		var cud models.Cud
		cud.Model = "post"
		cud.Command = "delete"
		cud.Post = *mp
		createQueue.QueueCreateCache(cud)
		log.Println("----------QUEUE CACHE DELETE POST SENDED-------------")
	}

	return nil
}

func (w *WriteServ) GetPost(id string) (*models.Post, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("service: couldn't parse id")
	}

	read, err := w.postgresDB.GetPostRead(id)
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
