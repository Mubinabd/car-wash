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

	pb "github.com/Mubinabd/car-wash/genproto"
)

type CartManager struct {
	collec *mongo.Collection
}

func NewCartManager(db *mongo.Database) *CartManager {
	return &CartManager{
		collec: db.Collection("cart"),
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

	var cart pb.Cart
	collection := c.collec.Database().Collection("cart")
	filter := bson.M{"id": req.Id}

	err := collection.FindOne(context.TODO(), filter).Decode(&cart)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no booking found with the given ID")
		}
		return nil, err
	}
	logger.Info("Cart retrieved successfully", logrus.Fields{
		"cart_id": cart.Id,
	})
	return &cart, nil
}
