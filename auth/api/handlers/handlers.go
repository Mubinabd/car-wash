package handlers

import (
	"github.com/go-redis/redis/v8"
	kafka "github.com/Mubinabd/car-wash/pkg/kafka/producer"
	"github.com/Mubinabd/car-wash/pkg/logger"
	"github.com/Mubinabd/car-wash/service"
)

type Handlers struct {
	Auth     *service.AuthService
	User     *service.UserService
	RDB      *redis.Client
	Producer kafka.KafkaProducer
	Logger   *logger.Logger
}

func NewHandler(auth *service.AuthService, user *service.UserService, rdb *redis.Client, pr *kafka.KafkaProducer, l *logger.Logger) *Handlers {
	return &Handlers{
		Auth:     auth,
		User:     user,
		RDB:      rdb,
		Producer: *pr,
		Logger:   l,
	}
}
