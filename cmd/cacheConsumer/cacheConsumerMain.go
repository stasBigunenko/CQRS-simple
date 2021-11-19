package main

import (
	"CQRS-simple/cmd/http/myConfig"
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/cacheConsumer"
	"CQRS-simple/pkg/storage/redis"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func main() {
	config := myConfig.SetConfig()

	path := os.Getenv("RMQ_PATH")
	if path == "" {
		path = "localhost:5672/"
	}

	login := os.Getenv("RMQ_LOG")
	if login == "" {
		login = "guest"
	}

	pass := os.Getenv("RMQ_PASS")
	if pass == "" {
		pass = "guest"
	}

	connStr := "amqp://" + login + ":" + pass + "@" + path
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
		os.Exit(3)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		"read",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if msgs == nil {
		os.Exit(3)
	}

	// create connection with redis storage
	redisDB := redis.NewRedisDB(config.RedisAddr, config.RedisDB)

	cacheConsumer := cacheConsumer.NewCacheConsumer(redisDB)

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			var cud models.Cud
			json.Unmarshal(m.Body, &cud)
			log.Printf("----------QUEUE CACHE RECEIVED with FOLOWING DATA %v-------------\n", cud)
			cacheConsumer.Received(cud)
		}
	}()

	log.Println("Successfully Connected to RabbitMQ Instance")
	log.Println(" [*] - Waiting for messages")
	<-forever
}
