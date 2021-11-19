package cacheConsumer

import (
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/storage/redis"
)

type CacheConsumer struct {
	dbRedis redis.RedisDBInterface
}

func NewCacheConsumer(dbr redis.RedisDBInterface) *CacheConsumer {
	return &CacheConsumer{
		dbRedis: dbr,
	}
}

func (c *CacheConsumer) Received(cud models.Cud) {
	switch cud.Model {
	case "user":
		c.ReceivedUser(cud)
	case "post":
		c.ReceivedPost(cud)
	}
}

func (c *CacheConsumer) ReceivedUser(cud models.Cud) {
	switch cud.Command {
	case "update":
		c.UpdateUser(cud.User)
	case "delete":
		c.DeleteUser(cud.User)
	}
}

func (c *CacheConsumer) ReceivedPost(cud models.Cud) {
	switch cud.Command {
	case "create":
		c.CreatePost(cud.Post)
	case "update":
		c.UpdatePost(cud.Post)
	case "delete":
		c.DeletePost(cud.Post)
	}
}

func (c *CacheConsumer) UpdateUser(u models.User) {
	c.dbRedis.UpdateUser(u)
}

func (c *CacheConsumer) DeleteUser(u models.User) {
	c.dbRedis.DeleteUser(u.ID)
}

func (c *CacheConsumer) CreatePost(p models.Post) {
	c.dbRedis.CreatePost(p)
}

func (c *CacheConsumer) UpdatePost(p models.Post) {
	c.dbRedis.UpdatePost(p)
}

func (c *CacheConsumer) DeletePost(p models.Post) {
	c.dbRedis.DeletePost(p.ID, p.UserID)
}
