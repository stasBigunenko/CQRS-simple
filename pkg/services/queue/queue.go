package queue

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/inMemory"
	"CQRS-simple/pkg/storage/postgreSQL"
)

type Queue struct {
	queue   postgreSQL.DBInterface
	storage inMemory.InMemoryInterface
}

func NewQueue(q postgreSQL.DBInterface, s inMemory.InMemoryInterface) Queue {
	return Queue{
		queue:   q,
		storage: s,
	}
}

func (q *Queue) GetAllUsers() (*[]models.User, error) {

	users, err := q.queue.GetAllUsers()
	if err != nil {
		return &[]models.User{}, err
	}

	for _, val := range *users {
		var r models.Read
		r.User = val
		q.storage.CreateUser(r)
	}
	//users, err := q.storage.GetAllUsers()
	//if err != nil {
	//	return &[]models.User{}, err
	//}

	return users, nil
}

func (q *Queue) UserPosts(userID string) (*models.UserPosts, error) {
	postRead, err := q.storage.GetUserPosts(userID)
	if err != nil {
		return &models.UserPosts{}, err
	}

	return postRead, nil
}

//functions for inMemory storage
//func (q *Queue) Get(id string) (*models.Read, error) {
//
//	_, err := uuid.Parse(id)
//	if err != nil {
//		return nil, errors.New("service: couldn't parse id")
//	}
//
//	read, err := q.queue.Get(id)
//	if err != nil {
//		return nil, err
//	}
//
//	return &read, nil
//}

//func (q *Queue) GetUser(id string) (*models.User, error) {
//
//	_, err := uuid.Parse(id)
//	if err != nil {
//		return nil, errors.New("service: couldn't parse id")
//	}
//
//	read, err := q.queue.GetUserRead(id)
//	if err != nil {
//		return nil, err
//	}
//
//	user := models.User{
//		ID:   read.User.ID,
//		Name: read.User.Name,
//		Age:  read.User.Age,
//	}
//
//	return &user, nil
//}
//
//func (q *Queue) GetPost(id string) (*models.Post, error) {
//
//	_, err := uuid.Parse(id)
//	if err != nil {
//		return nil, errors.New("service: couldn't parse id")
//	}
//
//	read, err := q.queue.GetPostRead(id)
//	if err != nil {
//		return nil, err
//	}
//
//	post := models.Post{
//		ID:      read.PostRead.ID,
//		UserID:  read.User.ID,
//		Title:   read.PostRead.Title,
//		Message: read.PostRead.Message,
//	}
//
//	return &post, nil
//}
//
//func (q *Queue) UserPosts(userID string) (*models.UserPosts, error) {
//
//	userRead, err := q.queue.GetUserRead(userID)
//	if err != nil {
//		return &models.UserPosts{}, err
//	}
//
//	user := models.User{
//		ID:   userRead.User.ID,
//		Name: userRead.User.Name,
//		Age:  userRead.User.Age,
//	}
//
//	postsRead, err := q.queue.GetPosts(userID)
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

//func (q *Queue) GetAllUsers() (*[]models.User, error) {
//
//	users, err := q.queue.GetAllUsers()
//	if err != nil {
//		return &[]models.User{}, err
//	}
//
//	return users, nil
//}
