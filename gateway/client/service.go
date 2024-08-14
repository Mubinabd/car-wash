package client

import (
	"log"
	"log/slog"

	pbc "github.com/Mubinabd/car-wash/genproto"
	kafka "github.com/Mubinabd/car-wash/kafka/producer"
	cfg "github.com/Mubinabd/car-wash/load"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	ProviderClient pbc.ProviderServiceClient
	BookingClient  pbc.BookingsClient
	Service        pbc.ServicesServiceClient
	Notification   pbc.NotificationServiceClient
	Cart           pbc.CartServiceClient
	Reviews        pbc.ReviewServiceClient
	Payments       pbc.PaymentServiceClient
	KafkaProducer  kafka.KafkaProducer
}

func NewClients(cfg *cfg.Config) (*Clients, error) {
	slog.Info("new client")
	conn, err := grpc.Dial("mongo-db:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	kafkaProducer, err := kafka.NewKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		return nil, err
	}

	clients := &Clients{
		ProviderClient: pbc.NewProviderServiceClient(conn),
		BookingClient:  pbc.NewBookingsClient(conn),
		Service:        pbc.NewServicesServiceClient(conn),
		Notification:   pbc.NewNotificationServiceClient(conn),
		Cart:           pbc.NewCartServiceClient(conn),
		Reviews:        pbc.NewReviewServiceClient(conn),
		Payments:       pbc.NewPaymentServiceClient(conn),
		KafkaProducer:  kafkaProducer,
	}

	return clients, nil
}
