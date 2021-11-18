package redis

import (
	"CQRS-simple/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)

type RedisDB struct {
	Client *redis.Client
	mu     sync.Mutex
}

func NewRedisDB(addr string, db string) *RedisDB {

	rdb, _ := strconv.Atoi(db)

	redisDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       rdb,
	})

	val, err := redisDB.Ping().Result()
	fmt.Println(val, err)

	return &RedisDB{Client: redisDB}
}

func (r *RedisDB) CreateUser(ur models.Read) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	read := models.UserPosts{
		User: ur.User,
	}

	jr, err := json.Marshal(read)
	if err != nil {
		return errors.New("redis marshal problem")
	}

	err = r.Client.Set(read.User.ID, jr, 15*time.Second).Err()
	if err != nil {
		return errors.New("redis internal problem")
	}
	return nil
}
func (r *RedisDB) CreatePost(ur models.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	val, err := r.Client.Get(ur.UserID).Bytes()
	if err != nil {
		return errors.New("redis internal problem")
	}

	jur := models.UserPosts{}

	err = json.Unmarshal(val, &jur)
	if err != nil {
		return errors.New("redis unmarshal problems")
	}

	var pr models.PostRead

	pr.ID = ur.ID
	pr.Title = ur.Title
	pr.Message = ur.Message

	jur.Posts = append(jur.Posts, pr)

	jr, err := json.Marshal(jur)
	if err != nil {
		return errors.New("redis marshal problem")
	}

	err = r.Client.Set(ur.UserID, jr, 15*time.Second).Err()
	if err != nil {
		return errors.New("redis internal problem")
	}

	return nil
}
func (r *RedisDB) GetAllUsers() (*[]models.User, error) {
	var users []models.User

	all, err := r.Client.Keys("*").Result()
	if err != nil {
		return nil, errors.New("redis internal problem")
	}

	for _, val := range all {
		res, err := r.Client.Get(val).Bytes()
		if err != nil {
			return nil, errors.New("redis internal problem")
		}
		ju := models.UserPosts{}

		err = json.Unmarshal(res, &ju)
		if err != nil {
			return nil, errors.New("redis internal problem")
		}
		users = append(users, ju.User)
	}

	return &users, nil
}
func (r *RedisDB) GetUserPosts(id string) (*models.UserPosts, error) {
	var userPost models.UserPosts

	userPostsJ, err := r.Client.Get(id).Bytes()
	if err != nil {
		return nil, errors.New("redis internal problem")
	}

	err = json.Unmarshal(userPostsJ, &userPost)

	return &userPost, nil
}
func (r *RedisDB) UpdateUser(u models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	jUserPost, err := r.Client.Get(u.ID).Bytes()
	if err != nil {
		return errors.New("redis: user not found")
	}

	var userPost models.UserPosts

	err = json.Unmarshal(jUserPost, &userPost)
	if err != nil {
		return errors.New("redis: marshal problems")
	}

	userPost.User = u

	ju, err := json.Marshal(userPost)
	if err != nil {
		return errors.New("redis marshal problem")
	}

	err = r.Client.Set(u.ID, ju, 15*time.Second).Err()
	if err != nil {
		return errors.New("redis internal problem")
	}

	return nil
}

func (r *RedisDB) UpdatePost(p models.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	jUserPost, err := r.Client.Get(p.UserID).Bytes()
	if err != nil {
		return errors.New("redis: post not found")
	}

	var userPost models.UserPosts

	err = json.Unmarshal(jUserPost, &userPost)
	if err != nil {
		return errors.New("redis: post not found")
	}

	for i, val := range userPost.Posts {
		if val.ID == p.ID {
			userPost.Posts[i].Title = p.Title
			userPost.Posts[i].Message = p.Message
			break
		}
	}

	ju, err := json.Marshal(userPost)
	if err != nil {
		return errors.New("redis marshal problem")
	}

	err = r.Client.Set(p.UserID, ju, 15*time.Second).Err()
	if err != nil {
		return errors.New("redis internal problem")
	}

	return nil
}
func (r *RedisDB) DeleteUser(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	err := r.Client.Del(id).Err()
	if err != nil {
		return errors.New("redis internal problem")
	}
	return nil
}
func (r *RedisDB) DeletePost(id, userID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	jUserPost, err := r.Client.Get(userID).Bytes()
	if err != nil {
		return errors.New("redis: post not found")
	}

	var userPost models.UserPosts

	err = json.Unmarshal(jUserPost, &userPost)
	if err != nil {
		return errors.New("redis: post not found")
	}

	for i, val := range userPost.Posts {
		if val.ID == id {
			if len(userPost.Posts)-1 == 0 {
				userPost.Posts = userPost.Posts[:0]
				break
			} else if i < len(userPost.Posts)-1 {
				userPost.Posts = append(userPost.Posts[:i], userPost.Posts[i+1:]...)
				break
			} else {
				userPost.Posts = append(userPost.Posts[:i])
				break
			}
		}
	}

	ju, err := json.Marshal(userPost)
	if err != nil {
		return errors.New("redis marshal problem")
	}

	err = r.Client.Set(userID, ju, 15*time.Second).Err()
	if err != nil {
		return errors.New("redis internal problem")
	}

	return nil
}

func (r *RedisDB) Exist(id string) bool {
	_, err := r.Client.Get(id).Bytes()
	if err != nil {
		return false
	}
	return true
}
