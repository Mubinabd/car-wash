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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServicesManager struct {
	collec *mongo.Collection
}

func NewServicesManager(db *mongo.Database) *ServicesManager {
	return &ServicesManager{
		collec: db.Collection("services"),
	}
}

func (s *ServicesManager) AddService(req *pb.AddServiceReq) (*pb.Empty, error) {
	res, err := s.collec.InsertOne(context.Background(), req)
	if err != nil {
		log.Println("error while inserting", err)
		return nil, err
	}

	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	logger.Info("Service created successfully")

	return &pb.Empty{}, nil
}

func (s *ServicesManager) GetService(req *pb.GetById) (*pb.Services, error) {
	var service pb.Services
	collection := s.collec
	filter := bson.M{"id": req.Id}

	err := collection.FindOne(context.TODO(), filter).Decode(&service)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no service found with the given ID")
		}
		return nil, err
	}

	logger.Info("Service retrieved successfully", logrus.Fields{
		"service_id":   service.Id,
		"service_name": service.Name,
	})

	return &service, nil
}

func (s *ServicesManager) ListAllServices(req *pb.ListAllServicesReq) (*pb.ListAllServicesResp, error) {
	var services []*pb.Services
	collection := s.collec.Database().Collection("services")
	filter := bson.M{}

	if req.Name != "" {
		filter["name"] = req.Name
	}
	if req.Description != "" {
		filter["description"] = req.Description
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
		var service pb.Services
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	logger.Info("Services listed successfully", logrus.Fields{
		"count": len(services),
	})

	return &pb.ListAllServicesResp{Services: services}, nil
}

func (s *ServicesManager) UpdateService(req *pb.UpdateServiceReq) (*pb.UpdateServiceResp, error) {
	filter := bson.M{"id": req.Id}

	update := bson.M{"$set": req.Services}

	result, err := s.collec.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	var updatedService pb.Services
	err = s.collec.FindOne(context.TODO(), filter).Decode(&updatedService)
	if err != nil {
		return nil, err
	}

	logger.Info("Service updated successfully", logrus.Fields{
		"count": result.ModifiedCount,
	})

	return &pb.UpdateServiceResp{Services: &updatedService}, nil
}
func (s *ServicesManager) DeleteService(req *pb.DeleteServiesReq) (*pb.DeleteServiesResp, error) {
	filter := bson.M{"id": req.Id}
	_, err := s.collec.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	logger.Info("Services deleted successfully", logrus.Fields{
		"country_id": req.Id,
	})

	return &pb.DeleteServiesResp{Success: true, Message: "Services deleted successfully"}, nil
}

func (s *ServicesManager) SearchServices(req *pb.SearchServicessReq) (*pb.SearchServicessResp, error) {
	var services []*pb.Services

	collection := s.collec.Database().Collection("services")

	filter := bson.M{}
	if req.Name != "" {
		filter["name"] = req.Name
	}
	if req.Description != "" {
		filter["description"] = req.Description
	}
	if req.Price != 0 {
		filter["price"] = req.Price
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var service pb.Services
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	logger.Info("Services listed successfully", logrus.Fields{
		"count": len(services),
	})

	return &pb.SearchServicessResp{Services: services}, nil
}

func (s *ServicesManager) GetServicesByPriceRange(req *pb.GetServicesByPriceRangeReq) (*pb.GetServicesByPriceRangeResp, error) {
	var services []*pb.Services
	collection := s.collec.Database().Collection("services")

	filter := bson.M{}
	if req.MinPrice > 0 {
		filter["price"] = bson.M{"$gte": req.MinPrice}
	}
	if req.MaxPrice > 0 {
		if filter["price"] == nil {
			filter["price"] = bson.M{"$lte": req.MaxPrice}
		} else {
			filter["price"].(bson.M)["$lte"] = req.MaxPrice
		}
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var service pb.Services
		if err := cursor.Decode(&service); err != nil {
			return nil, err
		}
		services = append(services, &service)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	logger.Info("Services listed successfully", logrus.Fields{
		"count": len(services),
	})

	return &pb.GetServicesByPriceRangeResp{Services: services}, nil
}
