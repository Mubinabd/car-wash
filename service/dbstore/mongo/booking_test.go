package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

func (m *MockNotificationManager) AddNotification(req *pb.AddNotificationReq) (*pb.Empty, error) {
	args := m.Called(req)
	return args.Get(0).(*pb.Empty), args.Error(1)
}

func TestAddBooking(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	notificationManager := NewNotificationManager(db)
	bookingManager := NewBookingManager(db, notificationManager)

	req := &pb.AddBookingReq{
		UserId:       "user123",
		ProviderId:   "provider456",
		ServiceId:    "service789",
		SchuduleTime: time.Now().String(),
	}

	_, err = bookingManager.AddBooking(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
func TestGetBooking(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	notificationManager := NewNotificationManager(db)
	bookingManager := NewBookingManager(db, notificationManager)

	bookingID := primitive.NewObjectID()
	booking := &pb.Booking{
		Id:           bookingID.Hex(),
		UserId:       "user123",
		ProviderId:   "provider456",
		ServiceId:    "service789",
		SchuduleTime: time.Now().String(),
		Status:       "confirmed",
		TotalPrice:   100.0,
	}
	_, err = db.Collection("bookings").InsertOne(context.Background(), booking)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.GetById{Id: bookingID.Hex()}
	resp, err := bookingManager.GetBooking(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, bookingID.Hex(), resp.Booking.Id, "Booking ID should match")
}

func TestDeleteBooking(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	notificationManager := NewNotificationManager(db)
	bookingManager := NewBookingManager(db, notificationManager)

	bookingID := primitive.NewObjectID()
	booking := &pb.Booking{
		Id:           bookingID.Hex(),
		UserId:       "user123",
		ProviderId:   "provider456",
		ServiceId:    "service789",
		SchuduleTime: time.Now().String(),
		Status:       "confirmed",
		TotalPrice:   100.0,
	}
	_, err = db.Collection("bookings").InsertOne(context.Background(), booking)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.DeleteBookingReq{Id: bookingID.Hex()}
	_, err = bookingManager.DeleteBooking(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = db.Collection("bookings").FindOne(context.Background(), bson.M{"_id": bookingID}).Err()
	assert.Equal(t, mongo.ErrNoDocuments, err, "Expected booking to be deleted")
}
