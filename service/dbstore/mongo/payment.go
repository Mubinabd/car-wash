package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Mubinabd/car-wash/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

type PaymentsManager struct {
	collec       *mongo.Collection
	notification *NotificationManager
}

func NewPaymentsManager(db *mongo.Database, notificationManager *NotificationManager) *PaymentsManager {
	return &PaymentsManager{
		collec:       db.Collection("payments"),
		notification: notificationManager,
	}
}
func (p *PaymentsManager) AddPayment(req *pb.AddPaymentReq) (*pb.Empty, error) {
	ctx := context.Background()

	_, err := p.collec.InsertOne(ctx, req)
	if err != nil {
		return nil, err
	}

	notificationReq := &pb.AddNotificationReq{
		UserId:  req.BookingId,
		Message: "Your payment has been successfully created!",
	}

	cartCollection := p.collec.Database().Collection("cart")
	_, err = cartCollection.UpdateOne(
		ctx,
		bson.M{"_id": req.CartId},
		bson.M{"$inc": bson.M{"total": -req.Amount}},
	)
	if err != nil {
		return nil, err
	}
	_, err = p.notification.AddNotification(notificationReq)
	if err != nil {
		return nil, err
	}

	logger.Info("Payment created and cart updated successfully")

	return &pb.Empty{}, nil
}

func (p *PaymentsManager) GetPayment(req *pb.GetById) (*pb.GetPaymentResp, error) {
	var payment pb.Payment
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}
	err = p.collec.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no payment found with the given ID")
		}
		return nil, err
	}

	logger.Info("Payment retrieved successfully", logrus.Fields{
		"payment_id":     payment.Id,
		"payment_amount": payment.Amount,
	})

	return &pb.GetPaymentResp{Payment: &payment}, nil
}

func (p *PaymentsManager) ListAllPayments(req *pb.ListAllPaymentsReq) (*pb.ListAllPaymentsResp, error) {

	var payments []*pb.Payment
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

	cursor, err := p.collec.Find(context.TODO(), filter,
		options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var payment pb.Payment
		if err := cursor.Decode(&payment); err != nil {
			return nil, err
		}
		payments = append(payments, &payment)
	}

	logger.Info("Payments retrieved successfully", logrus.Fields{
		"payments_count": len(payments),
	})

	return &pb.ListAllPaymentsResp{Payments: payments}, nil
}
