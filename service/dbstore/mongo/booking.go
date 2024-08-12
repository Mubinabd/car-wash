package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Mubinabd/car-wash/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

type BookingManager struct {
	collec             *mongo.Collection
	notificationManager *NotificationManager
}

func NewBookingManager(collec *mongo.Database, notificationManager *NotificationManager) *BookingManager {
	return &BookingManager{
		collec: collec.Collection("bookings"),
		notificationManager: notificationManager,
	}
}


func (r *BookingManager) AddBooking(req *pb.AddBookingReq) (*pb.Empty, error) {
	ctx := context.Background()
	session, err := r.collec.Database().Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	transactionOperation := func(sessionContext mongo.SessionContext) (interface{}, error) {
		res, err := r.collec.InsertOne(sessionContext, req)
		if err != nil {
			return nil, err
		}

		if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
			return nil, errors.New("failed to convert InsertedID to ObjectID")
		}

		notificationReq := &pb.AddNotificationReq{
			UserId:   req.UserId,
			Message:  "Your booking has been successfully created!",
			BookingId: res.InsertedID.(primitive.ObjectID).Hex(),
		}

		_, err = r.notificationManager.AddNotification(notificationReq)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
	_, err = session.WithTransaction(ctx, transactionOperation)
	if err != nil {
		log.Println("Transaction failed:", err)
		return nil, err
	}

	logger.Info("Booking and notification created successfully")

	return &pb.Empty{}, nil
}

func (r *BookingManager) GetBooking(req *pb.GetById) (*pb.Booking, error) {
	var booking pb.Booking
	collection := r.collec.Database().Collection("bookings")
	filter := bson.M{"id": req.Id}

	err := collection.FindOne(context.TODO(), filter).Decode(&booking)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no booking found with the given ID")
		}
		return nil, err
	}
	logger.Info("Booking retrieved successfully", logrus.Fields{
		"booking_id": booking.Id,
	})
	return &booking, nil
}

func (r *BookingManager) ListAllBookings(req *pb.ListAllBookingsReq) (*pb.ListAllBookingsResp, error) {
	var bookings []*pb.Booking
	collection := r.collec.Database().Collection("bookings")
	filter := bson.M{}

	if req.Status != "" {
		filter["status"] = req.Status
	}
	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}
	if req.ProviderId != "" {
		filter["provider_id"] = req.ProviderId
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

	cursor, err := collection.Find(context.TODO(), len(filter),
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
	return &pb.ListAllBookingsResp{Bookings: bookings}, nil
}

func (r *BookingManager) UpdateBooking(req *pb.UpdateBookingReq) (*pb.UpdateBookingResp, error) {
	filter := bson.M{"id": req.Id}
	update := bson.M{"$set": req}

	result, err := r.collec.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var updatedBooking pb.Booking
	err = r.collec.FindOne(context.TODO(), filter).Decode(&updatedBooking)
	if err != nil {
		return nil, err
	}
	logger.Info("Booking updated successfully", logrus.Fields{
		"count": result.ModifiedCount,
	})
	return &pb.UpdateBookingResp{Success: true,Message: "Booking updated successfully"}, nil
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

func(r *BookingManager) GetBookingsByProvider(req *pb.BookingsByProviderReq) (*pb.BookingsByProviderResp, error) {

	var bookings []*pb.Booking
	collection := r.collec
	filter := bson.M{"provider_id": req.ProviderId}

	cursor, err := collection.Find(context.TODO(), filter)
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

	return &pb.BookingsByProviderResp{Bookings: bookings}, nil	

}