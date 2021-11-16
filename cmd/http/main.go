package main

import (
	"CQRS-simple/cmd/http/myConfig"
	"CQRS-simple/pkg/handlers"
	"CQRS-simple/pkg/services/command"
	"CQRS-simple/pkg/services/queue"
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

	//storage := inMemory.NewInMemory()

	storage := redis.NewRedisDB(config.RedisAddr, config.RedisDB)

	db, err := postgreSQL.NewPDB(config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPsw, config.PostgresDB, config.PostgresSSL)
	if err != nil {
		if err != nil {
			log.Fatalf("failed to connect postgreSQL: %s", err)
		}
	}

	command := command.NewCommand(db, storage)
	queu := queue.NewQueue(db, storage)
	userRoutes := handlers.NewHandler(&command, &queu)

	r := mux.NewRouter()

	router := userRoutes.Routes(r)

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
