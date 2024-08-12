package mongo

import (
	"context"
	"fmt"
	"log"

	"github.com/Mubinabd/car-wash/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/Mubinabd/car-wash/genproto"
)

type PaymentsManager struct {
	collec *mongo.Collection
}

func NewPaymentsManager(db *mongo.Database) *PaymentsManager {
	return &PaymentsManager{
		collec: db.Collection("payments"),
	}
}
func (p *PaymentsManager) AddPayment(req *pb.AddPaymentReq) (*pb.Empty, error) {
	ctx := context.Background()

	session, err := p.collec.Database().Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)

	transactionOperation := func(ctxs mongo.SessionContext) (interface{}, error) {
		_, err := p.collec.InsertOne(ctxs, req)
		if err != nil {
			return nil, err
		}

		cartCollection := p.collec.Database().Collection("cart")
		_, err = cartCollection.UpdateOne(
			ctxs,
			bson.M{"_id": req.CartId},
			bson.M{"$inc": bson.M{"total": -req.Amount}},
		)
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

	logger.Info("Payment created and cart updated successfully")

	return &pb.Empty{}, nil
}

func (p *PaymentsManager) GetPayment(req *pb.GetById) (*pb.GetPaymentResp, error) {
	var payment pb.GetPaymentResp
	collection := p.collec
	filter := bson.M{"id": req.Id}

	err := collection.FindOne(context.TODO(), filter).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no payment found with the given ID")
		}
		return nil, err
	}

	logger.Info("Payment retrieved successfully", logrus.Fields{
		"payment_id":     payment.Payment.Id,
		"payment_amount": payment.Payment.Amount,
	})

	return &payment, nil
}

func (p *PaymentsManager) ListAllPayments(req *pb.ListAllPaymentsReq) (*pb.ListAllPaymentsResp, error) {

	var payments []*pb.Payment
	collection := p.collec.Database().Collection("payments")
	filter := bson.M{}

	if req.BookingId != "" {
		filter["booking_id"] = req.BookingId
	}
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

	cursor, err := collection.Find(context.TODO(), filter,
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
