package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

func setupTestDB(t *testing.T) (*mongo.Database, *PaymentsManager) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { client.Disconnect(context.Background()) })

	db := client.Database("testdb")
	notificationManager := NewNotificationManager(db) 
	paymentsManager := NewPaymentsManager(db, notificationManager)
	return db, paymentsManager
}

func TestAddPayment(t *testing.T) {
	db, paymentsManager := setupTestDB(t)

	paymentID := primitive.NewObjectID()
	payment := &pb.AddPaymentReq{
		BookingId: "booking123",
		CartId:    "cart123",
		Amount:    50.0,
	}

	_, err := paymentsManager.AddPayment(payment)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result pb.Payment
	err = db.Collection("payments").FindOne(context.Background(), bson.M{"_id": paymentID}).Decode(&result)
	if err != nil {
		t.Fatalf("expected payment to be found, got %v", err)
	}

	assert.Equal(t, paymentID.Hex(), result.Id, "Payment ID should match")
	assert.Equal(t, payment.Amount, result.Amount, "Payment amount should match")
}

func TestGetPayment(t *testing.T) {
	db, paymentsManager := setupTestDB(t)

	paymentID := primitive.NewObjectID()
	payment := &pb.Payment{
		Id:        paymentID.Hex(),
		BookingId: "booking123",
		CartId:    "cart123",
		Amount:    50.0,
	}

	_, err := db.Collection("payments").InsertOne(context.Background(), payment)
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.GetById{Id: paymentID.Hex()}
	resp, err := paymentsManager.GetPayment(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.Equal(t, paymentID.Hex(), resp.Payment.Id, "Payment ID should match")
	assert.Equal(t, payment.Amount, resp.Payment.Amount, "Payment amount should match")
}


