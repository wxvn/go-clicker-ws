package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/wxvn/go-clicker-ws/internal/logger"
	"github.com/wxvn/go-clicker-ws/internal/mongodb"
	"github.com/wxvn/go-clicker-ws/internal/redis"
)

type Config struct {
	Addr string
	Port int

	AllowedOrigins []string

	MongoDB mongodb.Config
	Logger  logger.Logger
	Redis   redis.Config
}

func Load() Config {
	_ = godotenv.Load()

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}

	mongoPort, err := strconv.Atoi(os.Getenv("MONGO_PORT"))
	if err != nil {
		mongoPort = 27017
	}

	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		redisDB = 0
	}

	return Config{
		Addr: os.Getenv("APP_ADDR"),
		Port: port,

		AllowedOrigins: strings.Split(os.Getenv("ALLOWED_ORIGINS"), ","),

		MongoDB: mongodb.Config{
			Host:     os.Getenv("MONGO_HOST"),
			Port:     mongoPort,
			Database: os.Getenv("MONGO_DATABASE"),
			User:     os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
		},

		Logger: logger.Logger{
			Level: os.Getenv("LOG_LEVEL"),
		},

		Redis: redis.Config{
			Addr:     os.Getenv("REDIS_ADDR"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		},
	}
}
