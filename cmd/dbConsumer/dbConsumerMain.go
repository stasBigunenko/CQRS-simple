package main

import (
	"CQRS-simple/cmd/http/myConfig"
	"CQRS-simple/pkg/models"
	"CQRS-simple/pkg/rabbitMQ/dbConsumer"
	"CQRS-simple/pkg/services/writeServ"
	"CQRS-simple/pkg/storage/postgreSQL"
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
		"cud",
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
	storage := redis.NewRedisDB(config.RedisAddr, config.RedisDB)

	// create connection with postgreSQL storage
	db, err := postgreSQL.NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
	if err != nil {
		if err != nil {
			log.Fatalf("failed to connect postgreSQL: %s", err)
		}
	}

	// interface of write functions
	writeServ := writeServ.NewWriteServ(db, storage)
	// interface for read functions

	dbConsumer := dbConsumer.NewDBConsumer(&writeServ)

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			var cud models.Cud
			json.Unmarshal(m.Body, &cud)
			dbConsumer.Received(cud)
		}
	}()

	log.Println("Successfully Connected to RabbitMQ Instance")
	log.Println(" [*] - Waiting for messages")
	<-forever
}
