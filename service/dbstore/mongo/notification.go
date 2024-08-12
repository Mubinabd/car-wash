package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationManager struct {
	collec *mongo.Collection
}

func NewNotificationManager(db *mongo.Database) *NotificationManager {

	return &NotificationManager{
		collec: db.Collection("notifications"),
	}
}

func (n *NotificationManager) AddNotification(req *pb.AddNotificationReq) (*pb.Empty, error) {
	res, err := n.collec.InsertOne(context.Background(), req)
	if err != nil {
		log.Println("error while inserting", err)
		return nil, err
	}

	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	logger.Info("Notification created successfully")

	return &pb.Empty{}, nil

}

func (n *NotificationManager) GetNotifications(req *pb.GetNotificationsReq) (*pb.Notification, error) {
	var notification pb.Notification
	collection := n.collec.Database().Collection("notifications")

	filter := bson.M{"user_id": req.UserId}

	err := collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no notification found for user ID %s", req.UserId)
		}
		return nil, err
	}

	logger.Info("Notification retrieved successfully", logrus.Fields{
		"notification_id": notification.Id,
		"user_id":         req.UserId,
	})

	return &notification, nil
}

func (n *NotificationManager) MarkNotificationAsRead(req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error) {
    
    objectId, err := primitive.ObjectIDFromHex(req.Id)
    if err != nil {
        return nil, fmt.Errorf("invalid resource id format: %s", req.Id)
    }

    filter := bson.M{"_id": objectId}
    update := bson.M{"$set": bson.M{"completed": true}}

    result, err := n.collec.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return nil, err
    }

    if result.MatchedCount == 0 {
        return nil, fmt.Errorf("resource not found")
    }

    response := &pb.MarkNotificationAsReadResp{
        Message:  "Resource marked as complete",
		Success:  true,
    }

    return response, nil
}

