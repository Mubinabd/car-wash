package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Mubinabd/car-wash/config"
	"github.com/Mubinabd/car-wash/dbstore/mongo"
	"github.com/Mubinabd/car-wash/pkg"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
	"github.com/Mubinabd/car-wash/service"

	"github.com/Mubinabd/car-wash/kafka"

	"google.golang.org/grpc"
)

func main() {
	logger.InitLog()
	logger.Info("Configuration loaded successfully")

	db, err := mongo.ConnectMongo()
	if err != nil {
		log.Fatal("error while connecting to mongo: ", err)
	}
	liss, err := net.Listen("tcp",config.Load().HTTPPort)
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	var wg sync.WaitGroup
	wg.Add(1)
	brokers := []string{"kafka:9092"}

	kcm := kafka.NewKafkaConsumerManager()

	bservice := service.NewBookingService(db)
	cservice := service.NewCartService(db)
	prservice := service.NewProviderService(db)
	nservice := service.NewNotificationService(db)
	pservice := service.NewPaymentService(db)
	rservice := service.NewReviewService(db)
	sservice := service.NewServiceService(db)

	pkg.Reader(brokers, kcm, bservice, cservice, prservice,rservice, nservice, sservice)

	gServer := grpc.NewServer()

	pb.RegisterBookingsServer(gServer, bservice)
	pb.RegisterCartServiceServer(gServer, cservice)
	pb.RegisterNotificationServiceServer(gServer, nservice)
	pb.RegisterPaymentServiceServer(gServer, pservice)
	pb.RegisterProviderServiceServer(gServer, prservice)
	pb.RegisterReviewServiceServer(gServer, rservice)
	pb.RegisterServicesServiceServer(gServer, sservice)

	log.Println("Server started on port", cfg.HTTPPort)
	if err := gServer.Serve(liss); err != nil {
		log.Fatal(err)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		sig := <-sigChan
		logger.Info("Received signal:", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		gServer.GracefulStop()
		<-ctx.Done()
		logger.Info("Graceful shutdown complete.")
	}
}
