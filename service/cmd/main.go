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
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
	"github.com/Mubinabd/car-wash/service"
	"github.com/Mubinabd/car-wash/dbstore/mongo"

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
	liss, err := net.Listen("tcp", ":27017")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()


	var wg sync.WaitGroup
	wg.Add(1)
	brokers := []string{"localhost:9092"}

	kcm := kafka.NewKafkaConsumerManager()

	

	bservice := service.NewBookingService(db)
	cservice := service.NewCartService(db)
	prservice := service.NewProviderService(db)
	nservice := service.NewNotificationService(db)
	pservice := service.NewPaymentService(db)
	rservice := service.NewReviewService(db)
	sservice := service.NewServiceService(db)


	if err := kcm.RegisterConsumer(brokers, "cr-booking", "product", kafka.BookingHandler(bservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "up-booking", "product", kafka.UpdateHandler(bservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-booking", "product", kafka.DeleteBookingHandler(bservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-cart", "product", kafka.CartHandler(cservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-booking' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-booking': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-provider", "product", kafka.ProviderHandler(prservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-provider' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-provider': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "up-provider", "product", kafka.UpdateProviderHandler(prservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'up-provider' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-provider': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-provider", "product", kafka.DeleteProviderHandler(prservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'dl-provider' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-provider': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-review", "product", kafka.ReviewHandler(rservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-review' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "up-review", "product", kafka.UpdatereviewHandler(rservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'up-review' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-review", "product", kafka.DeletereviewHandler(rservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'dk-review' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-review': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "cr-service", "service", kafka.ServiceHandler(sservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'cr-service' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'cr-service': %v", err)

		}
	}
	if err := kcm.RegisterConsumer(brokers, "dl-service", "service", kafka.DeleteserviceHandler(sservice)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'dl-service' already exists")
		} else {
			log.Printf("Failed to register consumer for topic 'dl-service': %v", err)

		}
	}
	
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
