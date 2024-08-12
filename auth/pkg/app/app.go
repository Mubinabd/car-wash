package app

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/go-redis/redis/v8"
	"github.com/Mubinabd/car-wash/api"
	"github.com/Mubinabd/car-wash/api/handlers"
	"github.com/Mubinabd/car-wash/config"
	kafka "github.com/Mubinabd/car-wash/pkg/kafka/consumer"
	prd "github.com/Mubinabd/car-wash/pkg/kafka/producer"
	"github.com/Mubinabd/car-wash/pkg/logger"
	"github.com/Mubinabd/car-wash/pkg/storage/postgres"
	"github.com/Mubinabd/car-wash/service"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func Run(cfg *config.Config) {
	basepath = "/home/mubina/Desktop/exams/5-exam/auth_service"
	l := logger.NewLogger(basepath, cfg.LOG_PATH)

	// Postgres Connection
	db, err := postgres.NewPostgresStorage(cfg)
	if err != nil {
		l.ERROR.Printf("can't connect to db: %v", err)
	}
	defer db.Db.Close()
	l.INFO.Println("Connected to Postgres")

	// Redis Connection
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",  
		Password: "",                
		DB:       0,
	})
	
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		l.ERROR.Panicf("Failed to connect to Redis: %v", err)
	}
	l.INFO.Println("Connected to Redis")

	authService := service.NewAuthService(db)
	userService := service.NewUserService(db)

	// Kafka
	brokers := []string{"localhost:9092"}
	cm := kafka.NewKafkaConsumerManager()
	pr, err := prd.NewKafkaProducer(brokers)
	if err != nil {
		l.ERROR.Println("Failed to create Kafka producer:", err)
		return
	}

	Reader(brokers, cm, authService, userService, l)

	// HTTP Server
	h := handlers.NewHandler(authService, userService, rdb, &pr, l)

	router := api.Engine(h)
	router.SetTrustedProxies(nil)

	if err := router.Run(cfg.AUTH_PORT); err != nil {
		l.ERROR.Panicf("can't start server: %v", err)
	}

	l.INFO.Printf("REST server started on port %s", cfg.AUTH_PORT)
}
