package readServ

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/postgreSQL"
	"CQRS-simple/pkg/storage/redis"
	"log"
)

type ReadServ struct {
	postgresDB postgreSQL.DBInterface
	redisDB    redis.RedisDBInterface
}

func NewReadServ(p postgreSQL.DBInterface, s redis.RedisDBInterface) ReadServ {
	return ReadServ{
		postgresDB: p,
		redisDB:    s,
	}
}

func (r *ReadServ) GetAllUsers() (*[]models.User, error) {

	users, err := r.postgresDB.GetAllUsers()
	if err != nil {
		return &[]models.User{}, err
	}

	for _, val := range *users {
		var rr models.Read
		rr.User = val
		r.redisDB.CreateUser(rr)
		res, err := r.postgresDB.GetPosts(rr.User.ID)
		if err != nil {
			continue
		}
		for _, p := range res {
			var pr models.Post
			pr.ID = p.ID
			pr.UserID = rr.User.ID
			pr.Title = p.Title
			pr.Message = p.Message
			r.redisDB.CreatePost(pr)
		}
	}

	return users, nil
}

func (r *ReadServ) UserPosts(userID string) (models.UserPosts, error) {
	var postRead models.UserPosts
	postRead, err := r.redisDB.GetUserPosts(userID)
	log.Printf("postRead in Queue before if = %v\n", postRead)
	if err != nil {
		user, err := r.postgresDB.GetUser(userID)
		if err != nil {
			return models.UserPosts{}, err
		}
		postRead.User, _ = r.redisDB.CreateUser(user)
		res, _ := r.postgresDB.GetPosts(userID)
		for _, p := range res {
			var pr models.Post
			pr.ID = p.ID
			pr.UserID = userID
			pr.Title = p.Title
			pr.Message = p.Message
			post, _ := r.redisDB.CreatePost(pr)
			postRead.Posts = append(postRead.Posts, post)
		}
	}
	log.Printf("postRead in Queue after if = %v\n", postRead)
	return postRead, nil
}
