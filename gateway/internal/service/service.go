package client

import (
	"log"

	pbc "github.com/Mubinabd/car-wash/genproto"
	kafka "github.com/Mubinabd/car-wash/internal/pkg/kafka/producer"
	cfg "github.com/Mubinabd/car-wash/internal/pkg/load"
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
	connections    []*grpc.ClientConn
}

func NewClients(cfg *cfg.Config) (*Clients, error) {
	conns := make([]*grpc.ClientConn, 0)

	conn, err := grpc.NewClient("localhost:8085", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	kafkaProducer, err := kafka.NewKafkaProducer([]string{"localhost:9092"})
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
		connections:    conns,
	}

	return clients, nil
}

func (c *Clients) CloseConnections() {
	for _, conn := range c.connections {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close gRPC connection: %v", err)
		}
	}
}
