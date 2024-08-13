package mongo

import (
	"context"
	"errors"

	"github.com/Mubinabd/car-wash/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

type BookingManager struct {
	collec              *mongo.Collection
	notificationManager *NotificationManager
}

func NewBookingManager(collec *mongo.Database, notificationManager *NotificationManager) *BookingManager {
	return &BookingManager{
		collec:              collec.Collection("bookings"),
		notificationManager: notificationManager,
	}
}

func (r *BookingManager) AddBooking(req *pb.AddBookingReq) (*pb.Empty, error) {
	ctx := context.Background()

	res, err := r.collec.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}

	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	notificationReq := &pb.AddNotificationReq{
		UserId:  req.UserId,
		Message: "Your booking has been successfully created!",
	}

	_, err = r.notificationManager.AddNotification(notificationReq)
	if err != nil {
		return nil, err
	}

	logger.Info("Booking and notification created successfully")

	return &pb.Empty{}, nil
}

func (r *BookingManager) GetBooking(req *pb.GetById) (*pb.GetBookingResp, error) {
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var booking pb.Booking
	err = r.collec.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&booking)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("booking not found")
	} else if err != nil {
		return nil, err
	}

	return &pb.GetBookingResp{Booking: &booking}, nil
}

func (r *BookingManager) ListAllBookings(req *pb.ListAllBookingsReq) (*pb.ListAllBookingsResp, error) {
	var bookings []*pb.Booking
	filter := bson.M{}

	if req.Status != "" {
		filter["status"] = req.Status
	}
	
	limit := int64(10)
	offset := int64(0)

	if req.Filter != nil {
		if req.Filter.Limit > 0 {
			limit = int64(req.Filter.Limit)
		}
		if req.Filter.Offset >= 0 {
			offset = int64(req.Filter.Offset)
		}
	}

	cursor, err := r.collec.Find(context.TODO(), filter,
		options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var booking pb.Booking
		if err := cursor.Decode(&booking); err != nil {
			return nil, err
		}
		bookings = append(bookings, &booking)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListAllBookingsResp{Bookings: bookings}, nil
}

func (r *BookingManager) UpdateBooking(req *pb.UpdateBookingReq) (*pb.UpdateBookingResp, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"provider_id":   req.Booking.ProviderId,
			"service_id":    req.Booking.ServiceId,
			"location":      req.Booking.Location,
			"schudule_time": req.Booking.SchuduleTime,
			"status":        req.Booking.Status,
			"total_price":   req.Booking.TotalPrice,
			"updated_at":    req.Booking.UpdatedAt,
		},
	}

	result, err := r.collec.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("no matching document found to update")
	}

	var updatedBooking pb.Booking
	err = r.collec.FindOne(context.TODO(), filter).Decode(&updatedBooking)
	if err != nil {
		return nil, err
	}

	logger.Info("Booking updated successfully", logrus.Fields{
		"count": result.ModifiedCount,
	})

	return &pb.UpdateBookingResp{Success: true, Message: "Booking updated successfully"}, nil
}


func (r *BookingManager) DeleteBooking(req *pb.DeleteBookingReq) (*pb.DeleteBookingResp, error) {
	filter := bson.M{"id": req.Id}
	_, err := r.collec.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	logger.Info("Booking deleted successfully", logrus.Fields{
		"booking_id": req.Id,
	})
	return &pb.DeleteBookingResp{}, nil
}
