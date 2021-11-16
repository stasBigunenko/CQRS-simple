package myConfig

import "os"

type Config struct {
	PortHTTP string

	//postgres
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPsw  string
	PostgresDB   string
	PostgresSSL  string

	//redis
	RedisAddr string
	RedisDB   string
}

func SetConfig() *Config {
	var config Config

	config.PortHTTP = os.Getenv("PORT_HTTP")
	if config.PortHTTP == "" {
		config.PortHTTP = ":8080"
	}

	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	if config.PostgresHost == "" {
		config.PostgresHost = "postgres"
	}

	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	if config.PostgresPort == "" {
		config.PostgresPort = "5432"
	}

	config.PostgresUser = os.Getenv("POSTGRES_USER")
	if config.PostgresUser == "" {
		config.PostgresUser = "localhost"
	}

	config.PostgresPsw = os.Getenv("POSTGRES_PASSWORD")
	if config.PostgresPsw == "" {
		config.PostgresPsw = "qwerty"
	}

	config.PostgresDB = os.Getenv("POSTGRES_DATABASE")
	if config.PostgresDB == "" {
		config.PostgresDB = "postgres"
	}

	config.PostgresSSL = os.Getenv("POSTGRES_SSL")
	if config.PostgresSSL == "" {
		config.PostgresSSL = "disable"
	}

	config.RedisAddr = os.Getenv("REDIS_ADDR")
	if config.RedisAddr == "" {
		config.RedisAddr = "127.0.0.1:6379"
	}

	config.RedisDB = os.Getenv("REDIS_DB")
	if config.RedisDB == "" {
		config.RedisDB = "redisDB"
	}

	return &Config{
		PortHTTP: config.PortHTTP,

		PostgresHost: config.PostgresHost,
		PostgresPort: config.PostgresPort,
		PostgresUser: config.PostgresUser,
		PostgresPsw:  config.PostgresPsw,
		PostgresDB:   config.PostgresDB,
		PostgresSSL:  config.PostgresSSL,

		RedisAddr: config.RedisAddr,
		RedisDB:   config.RedisDB,
	}
}
