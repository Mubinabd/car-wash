package mongo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Mubinabd/car-wash/dbstore"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageMongo struct {
	mongo         *mongo.Database
	BookingS      dbstore.BookingI
	CartS         dbstore.CartI
	NotificationS dbstore.NotificationI
	PaymentS      dbstore.PaymentI
	ProviderS     dbstore.ProviderI
	ReviewS       dbstore.ReviewI
	ServiceS      dbstore.ServiceI
}

func ConnectMongo() (dbstore.Storage, error) {
	uri := fmt.Sprintf("mongodb://%s:%d", "mongo-db", 8050)
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	slog.Info("Connected to MongoDB")

	database := client.Database("carwash")
	slog.Info("Database accessed successfully")

	slog.Info("Test document inserted")

	providerManager := NewProviderManager(database)
	cartManager := NewCartManager(database)
	notificationManager := NewNotificationManager(database)
	paymentManager := NewPaymentsManager(database,notificationManager)
	bookingManager := NewBookingManager(database, notificationManager)
	reviewManager := NewReviewsManager(database)
	serviceManager := NewServicesManager(database)

	return &StorageMongo{
		mongo:         database,
		BookingS:      bookingManager,
		CartS:         cartManager,
		PaymentS:      paymentManager,
		NotificationS: notificationManager,
		ReviewS:       reviewManager,
		ProviderS:     providerManager,
		ServiceS:      serviceManager,
	}, nil
}

func (s *StorageMongo) Booking() dbstore.BookingI {
	if s.BookingS == nil {
		notificationManager := NewNotificationManager(s.mongo)
		s.BookingS = NewBookingManager(s.mongo, notificationManager)
	}
	return s.BookingS
}

func (s *StorageMongo) Cart() dbstore.CartI {
	if s.CartS == nil {
		s.CartS = NewCartManager(s.mongo)
	}
	return s.CartS
}

func (s *StorageMongo) Notification() dbstore.NotificationI {
	if s.NotificationS == nil {
		s.NotificationS = NewNotificationManager(s.mongo)
	}
	return s.NotificationS
}
func (s *StorageMongo) Payment() dbstore.PaymentI {
	if s.PaymentS == nil {
		s.PaymentS = NewPaymentsManager(s.mongo,NewNotificationManager(s.mongo))
	}
	return s.PaymentS
}
func (s *StorageMongo) Provider() dbstore.ProviderI {
	if s.ProviderS == nil {
		s.ProviderS = NewProviderManager(s.mongo)
	}
	return s.ProviderS
}
func (s *StorageMongo) Review() dbstore.ReviewI {
	if s.ReviewS == nil {
		s.ReviewS = NewReviewsManager(s.mongo)
	}
	return s.ReviewS
}
func (s *StorageMongo) Servicee() dbstore.ServiceI {
	if s.ServiceS == nil {
		s.ServiceS = NewServicesManager(s.mongo)
	}
	return s.ServiceS
}
