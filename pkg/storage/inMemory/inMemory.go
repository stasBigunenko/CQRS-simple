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

func (i *InMemory) UpdateUser(u models.User) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	up, ok := i.storage[u.ID]
	if !ok {
		return errors.New("user not found")
	}

	up.User = u

	i.storage[u.ID] = up

	return nil
}

func (i *InMemory) UpdatePost(p models.Post) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	up, ok := i.storage[p.UserID]
	if !ok {
		return errors.New("user with this post not found")
	}

	for i, val := range up.Posts {
		if val.ID == p.ID {
			up.Posts[i].Title = p.Title
			up.Posts[i].Message = p.Message
		}
	}

	i.storage[p.UserID] = up

	return nil
}

func (i *InMemory) DeleteUser(id string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	_, ok := i.storage[id]
	if !ok {
		return errors.New("post can't be deleted - Id not found")
	}

	delete(i.storage, id)

	return nil
}

func (i *InMemory) DeletePost(id, userID string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	up, ok := i.storage[userID]
	if !ok {
		return errors.New("post can't be deleted - user not found")
	}

	for i, val := range up.Posts {
		if val.ID == id {
			if len(up.Posts)-1 == 0 {
				up.Posts = up.Posts[:0]
				break
			} else if i < len(up.Posts)-1 {
				up.Posts = append(up.Posts[:i], up.Posts[i+1:]...)
				break
			} else {
				up.Posts = append(up.Posts[:i])
				break
			}
		}
	}

	i.storage[userID] = up

	return nil
}
