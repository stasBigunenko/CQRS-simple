package queue

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/inMemory"
	"CQRS-simple/pkg/storage/postgreSQL"
	"log"
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
		res, err := q.queue.GetPosts(r.User.ID)
		if err != nil {
			continue
		}
		for _, p := range res {
			var pr models.Post
			pr.ID = p.ID
			pr.UserID = r.User.ID
			pr.Title = p.Title
			pr.Message = p.Message
			q.storage.CreatePost(pr)
		}
	}

	return users, nil
}

func (q *Queue) UserPosts(userID string) (models.UserPosts, error) {
	var postRead models.UserPosts
	postRead, err := q.storage.GetUserPosts(userID)
	log.Printf("postRead in Queue before if = %v\n", postRead)
	if err != nil {
		user, err := q.queue.GetUser(userID)
		if err != nil {
			return models.UserPosts{}, err
		}
		postRead.User, _ = q.storage.CreateUser(user)
		res, _ := q.queue.GetPosts(userID)
		for _, p := range res {
			var pr models.Post
			pr.ID = p.ID
			pr.UserID = userID
			pr.Title = p.Title
			pr.Message = p.Message
			post, _ := q.storage.CreatePost(pr)
			postRead.Posts = append(postRead.Posts, post)
		}
	}
	log.Printf("postRead in Queue after if = %v\n", postRead)
	return postRead, nil
}
