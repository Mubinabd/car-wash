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
	DB            *mongo.Database
	BookingS      dbstore.BookingI
	CartS         dbstore.CartI
	NotificationS dbstore.NotificationI
	PaymentS      dbstore.PaymentI
	ProviderS     dbstore.ProviderI
	ReviewS       dbstore.ReviewI
	ServiceS      dbstore.ServiceI
}

func ConnectMongo() (dbstore.Storage, error) {

	uri := fmt.Sprintf("mongodb://%s:%d",
		"localhost",
		27017,
	)
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

	mongo := client.Database("carwash")
	fmt.Println(mongo)

	providerManager := NewProviderManager(mongo)
	cartManager := NewCartManager(mongo)
	notificationManager := NewNotificationManager(mongo)
	paymentManager := NewPaymentsManager(mongo)
	bookingManager := NewBookingManager(mongo, notificationManager)
	reviewManager := NewReviewsManager(mongo)
	serviceManager := NewServicesManager(mongo)

	return &StorageMongo{
		DB:            mongo,
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
		notificationManager := NewNotificationManager(s.DB)
		s.BookingS = NewBookingManager(s.DB, notificationManager)
	}
	return s.BookingS
}

func (s *StorageMongo) Cart() dbstore.CartI {
	if s.CartS == nil {
		s.CartS = NewCartManager(s.DB)
	}
	return s.CartS
}

func (s *StorageMongo) Notification() dbstore.NotificationI {
	if s.NotificationS == nil {
		s.NotificationS = NewNotificationManager(s.DB)
	}
	return s.NotificationS
}
func (s *StorageMongo) Payment() dbstore.PaymentI {
	if s.PaymentS == nil {
		s.PaymentS = NewPaymentsManager(s.DB)
	}
	return s.PaymentS
}
func (s *StorageMongo) Provider() dbstore.ProviderI {
	if s.ProviderS == nil {
		s.ProviderS = NewProviderManager(s.DB)
	}
	return s.ProviderS
}
func (s *StorageMongo) Review() dbstore.ReviewI {
	if s.ReviewS == nil {
		s.ReviewS = NewReviewsManager(s.DB)
	}
	return s.ReviewS
}
func (s *StorageMongo) Service() dbstore.ServiceI {
	if s.ServiceS == nil {
		s.ServiceS = NewServicesManager(s.DB)
	}
	return s.ServiceS
}
