package inMemory

import (
	"CQRS-simple/pkg/models"
	"errors"
	"sync"
)

type InMemory struct {
	mu      sync.Mutex
	storage map[string]models.UserPosts
}

func NewInMemory() *InMemory {
	return &InMemory{
		storage: make(map[string]models.UserPosts),
	}
}

func (i *InMemory) CreateUser(ur models.Read) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	read := models.UserPosts{
		User: ur.User,
	}
	i.storage[ur.User.ID] = read

	return nil
}

func (i *InMemory) CreatePost(ur models.Read) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	user := i.storage[ur.User.ID].User

	posts := i.storage[ur.User.ID].Posts
	posts = append(posts, ur.PostRead)

	read := models.UserPosts{
		User:  user,
		Posts: posts,
	}

	i.storage[ur.User.ID] = read

	return nil
}

func (i *InMemory) GetAllUsers() (*[]models.User, error) {
	var u []models.User

	for _, v := range i.storage {
		ur := models.User{
			ID:   v.User.ID,
			Name: v.User.Name,
			Age:  v.User.Age,
		}
		u = append(u, ur)
	}
	return &u, nil
}

func (i *InMemory) GetUserPosts(id string) (*models.UserPosts, error) {

	var up models.UserPosts
	up, ok := i.storage[id]
	if !ok {
		return &models.UserPosts{}, errors.New("user not found")
	}

	return &up, nil
}
