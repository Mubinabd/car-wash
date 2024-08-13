package mongo

import (
	"context"
	"errors"
	"log"

	"github.com/Mubinabd/car-wash/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	pb "github.com/Mubinabd/car-wash/genproto"
)

type CartManager struct {
	collec *mongo.Collection
}

func NewCartManager(collec *mongo.Database) *CartManager {
	return &CartManager{
		collec: collec.Collection("cart"),
	}
}

func (c *CartManager) CreateCart(req *pb.CreateCartReq) (*pb.Empty, error) {

	res, err := c.collec.InsertOne(context.Background(), req)
	if err != nil {
		log.Println("error while inserting", err)
		return nil, err
	}

	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	logger.Info("Cart created successfully")

	return &pb.Empty{}, nil
}

func (c *CartManager) GetCart(req *pb.GetById) (*pb.Cart, error) {

	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var cart pb.Cart
	err = c.collec.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&cart)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("booking not found")
	} else if err != nil {
		return nil, err
	}

	return &cart, nil
}
