package myConfig

import "os"

type Config struct {
	// Hhtp connection
	PortHTTP string

	// Postgres
	PostgresHost string
	PostgresPort string
	PostgresUser string
	PostgresPsw  string
	PostgresDB   string
	PostgresSSL  string

	// Redis
	RedisAddr string
	RedisDB   string

	// Handler producer
	RMQPath string
	RMQLog  string
	RMQPass string
}

func SetConfig() *Config {
	var config Config

	config.PortHTTP = os.Getenv("PORT_HTTP")
	if config.PortHTTP == "" {
		config.PortHTTP = ":8080"
	}

	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	if config.PostgresHost == "" {
		config.PostgresHost = "localhost"
	}

	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	if config.PostgresPort == "" {
		config.PostgresPort = "5432"
	}

	config.PostgresUser = os.Getenv("POSTGRES_USER")
	if config.PostgresUser == "" {
		config.PostgresUser = "postgres"
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

	config.RMQPath = os.Getenv("RMQ_PATH")
	if config.RMQPath == "" {
		config.RMQPath = "localhost:5672/"
	}

	config.RMQLog = os.Getenv("RMQ_LOG")
	if config.RMQLog == "" {
		config.RMQLog = "guest"
	}

	config.RMQPass = os.Getenv("RMQ_PASS")
	if config.RMQPass == "" {
		config.RMQPass = "guest"
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

		RMQPath: config.RMQPath,
		RMQLog:  config.RMQLog,
		RMQPass: config.RMQPass,
	}
}
