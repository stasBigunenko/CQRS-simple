package main

import (
	"CQRS-simple/cmd/http/myConfig"
	"CQRS-simple/pkg/handlers"
	createQueue2 "CQRS-simple/pkg/rabbitMQ/createQueue"
	"CQRS-simple/pkg/services/readServ"
	"CQRS-simple/pkg/storage/postgreSQL"
	"CQRS-simple/pkg/storage/redis"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	config := myConfig.SetConfig()

	// create connection with redis storage
	storage := redis.NewRedisDB(config.RedisAddr, config.RedisDB)

	// create connection with postgreSQL storage
	db, err := postgreSQL.NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
	if err != nil {
		if err != nil {
			log.Fatalf("failed to connect postgreSQL: %s", err)
		}
	}

	// interface for read functions
	readServ := readServ.NewReadServ(db, storage)

	createQueueWrite := createQueue2.CreateQueue{}

	// create handler
	userRoutes := handlers.NewHandler(&readServ, &createQueueWrite)

	r := mux.NewRouter()

	router := userRoutes.Routes(r)

	// http server config
	srv := http.Server{
		Addr:    config.PortHTTP,
		Handler: router,
	}

	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		srv.Shutdown(context.Background())
	}()

	log.Printf("HTTP server started on port: %v\n", config.PortHTTP)
	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", )

	if err := srv.ListenAndServe(); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
