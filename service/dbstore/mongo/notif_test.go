package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

// Mock Database
type MockNotificationManager struct {
	mock.Mock
}
func (m *MockNotificationManager) GetNotifications(req *pb.GetNotificationsReq) (*pb.GetNotificationsResp, error) {
	args := m.Called(req)
	return args.Get(0).(*pb.GetNotificationsResp), args.Error(1)
}

func (m *MockNotificationManager) MarkNotificationAsRead(req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error) {
	args := m.Called(req)
	return args.Get(0).(*pb.MarkNotificationAsReadResp), args.Error(1)
}

func TestAddNotification(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	manager := NewNotificationManager(db)

	req := &pb.AddNotificationReq{
		UserId:  "user123",
		Message: "You have a new notification!",
	}

	_, err = manager.AddNotification(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Verify notification was inserted
	var notification pb.Notification
	err = db.Collection("notifications").FindOne(context.Background(), bson.M{"user_id": req.UserId}).Decode(&notification)
	if err != nil {
		t.Fatalf("expected notification to be found, got %v", err)
	}
}

func TestMarkNotificationAsRead(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("testdb")
	manager := NewNotificationManager(db)

	objectId := primitive.NewObjectID()
	_, err = db.Collection("notifications").InsertOne(context.Background(), &pb.Notification{
		Id:      objectId.Hex(),
		UserId:  "user123",
		Message: "Notification to be marked as read",
	})
	if err != nil {
		t.Fatal(err)
	}

	req := &pb.MarkNotificationAsReadReq{Id: objectId.Hex()}
	resp, err := manager.MarkNotificationAsRead(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assert.True(t, resp.Success, "Expected success response")

	// Verify notification was updated
	var notification pb.Notification
	err = db.Collection("notifications").FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&notification)
	if err != nil {
		t.Fatalf("expected notification to be found, got %v", err)
	}

	assert.True(t, notification.IsRead, "Expected notification to be marked as completed")
}
