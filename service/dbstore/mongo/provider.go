package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Mubinabd/car-wash/logger"
)


type ProviderManager struct {
    col *mongo.Collection
}

func NewProviderManager(db *mongo.Database) *ProviderManager {
    return &ProviderManager{
        col: db.Collection("provider"),
    }
}

func (p *ProviderManager) RegisterProvider(req *pb.RegisterProviderReq) (*pb.Empty, error) {
	res, err := p.col.InsertOne(context.Background(), req)

    if err != nil {
        log.Println("Error while inserting provider:", err)
        return nil, fmt.Errorf("failed to insert provider: %w", err)
    }

	if _, ok := res.InsertedID.(primitive.ObjectID); !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}
	
    log.Println("Provider created successfully with ID:", res.InsertedID)

    return &pb.Empty{}, nil
}

func (p *ProviderManager) GetProvider(req *pb.GetById) (*pb.Provider, error) {
	var provider pb.Provider
	collection := p.col.Database().Collection("provider")
	filter := bson.M{"id": req.Id}

	err := collection.FindOne(context.TODO(), filter).Decode(&provider)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no provider found with the given ID")
		}
		return nil, err
	}
	logger.Info("Provider retrieved successfully", logrus.Fields{
		"provider_id":  provider.Id,
		"company_name": provider.CompanyName,
	})

	return &provider, nil
}

func (p *ProviderManager) ListAllProviders(req *pb.ListAllProvidersReq) (*pb.ListAllProvidersResp, error) {
	var providers []*pb.Provider
	collection := p.col.Database().Collection("provider")

	filter := bson.M{}
	if req.CompanyName != "" {
		filter["company_name"] = req.CompanyName
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
		var provider pb.Provider
		if err := cursor.Decode(&provider); err != nil {
			return nil, err
		}
		providers = append(providers, &provider)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	logger.Info("Providers listed successfully", logrus.Fields{
		"count": len(providers),
	})

	return &pb.ListAllProvidersResp{Providers: providers}, nil
}

func (p *ProviderManager) DeleteProvider(req *pb.DeleteProviderReq) (*pb.DeleteProviderResp, error) {
	filter := bson.M{"id": req.Id}
	_, err := p.col.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	logger.Info("Provider deleted successfully", logrus.Fields{
		"country_id": req.Id,
	})

	return &pb.DeleteProviderResp{Success: true, Message: "Provider deleted successfully"}, nil
}

func (p *ProviderManager) UpdateProvider(req *pb.UpdateProviderReq) (*pb.UpdateProviderResp, error) {
	filter := bson.M{"id": req.Id}

	update := bson.M{"$set": req.Provider}

	result, err := p.col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, fmt.Errorf("provider with ID %s not found", req.Id)
	}

	logger.Info("Provider updated successfully", logrus.Fields{
		"provider_id": req.Id,
	})

	var updatedProvider pb.Provider
	err = p.col.FindOne(context.TODO(), filter).Decode(&updatedProvider)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProviderResp{Provider: &updatedProvider}, nil
}

func (p *ProviderManager) SearchProviders(req *pb.SearchProvidersReq) (*pb.SearchProvidersResp, error) {

	var providers []*pb.Provider
	collection := p.col.Database().Collection("provider")

	filter := bson.M{}
	if req.CompanyName != "" {
		filter["company_name"] = req.CompanyName
	}
	if req.Description != "" {
		filter["description"] = req.Description
	}
	if req.UserId != "" {
		filter["user_id"] = req.UserId
	}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var provider pb.Provider
		if err := cursor.Decode(&provider); err != nil {
			return nil, err
		}
		providers = append(providers, &provider)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	logger.Info("Providers listed successfully", logrus.Fields{
		"count": len(providers),
	})

	return &pb.SearchProvidersResp{Providers: providers}, nil
}
