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

type ReviewsManager struct {
	collec *mongo.Collection
}

func NewReviewsManager(db *mongo.Database) *ReviewsManager {
	return &ReviewsManager{
		collec: db.Collection("reviews"),
	}
}

func (r *ReviewsManager) AddReview(req *pb.AddReviewReq) (*pb.Empty, error) {
	res, err := r.collec.InsertOne(context.TODO(), req)
	if err != nil {
		return nil, err
	}

	_, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	logger.Info("Review created successfully")
	return &pb.Empty{}, nil
}

func (r *ReviewsManager) GetReview(req *pb.GetById) (*pb.Review, error) {
	if req.Id == "" {
        return nil, errors.New("id cannot be empty")
    }

    oid, err := primitive.ObjectIDFromHex(req.Id)
    if err != nil {
        return nil, errors.New("invalid ID format")
    }

    var review pb.Review
    err = r.collec.FindOne(context.TODO(), bson.M{"_id": oid}).Decode(&review)
    if err == mongo.ErrNoDocuments {
        return nil, errors.New("review not found")
    } else if err != nil {
        return nil, err
    }
	return &review, nil
}


func(r *ReviewsManager) UpdateReview(req *pb.UpdateReviewsReq) (*pb.UpdateReviewsResp, error) {
	filter := bson.M{"id": req.Id}

	update := bson.M{"$set": req.Review}	

	result, err := r.collec.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {		
		return nil, fmt.Errorf("review with ID %s not found", req.Id)
	}

	logger.Info("Review updated successfully", logrus.Fields{
		"review_id": req.Id,
	})
	return &pb.UpdateReviewsResp{Success: true,Message: "reviews updated successfully"}, nil
}


func(r *ReviewsManager) DeleteReview(req *pb.DeleteReviewReq) (*pb.DeleteReviewResp, error) {
	filter := bson.M{"id": req.Id}
	_, err := r.collec.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}		

	logger.Info("Review deleted successfully", logrus.Fields{
		"review_id": req.Id,
	})
	return &pb.DeleteReviewResp{Success: true,Message: "Review deleted successfully"}, nil
}
func(r *ReviewsManager)ListAllReviews(req *pb.ListAllReviewsReq) (*pb.ListAllReviewsResp, error) {

	var reviews []*pb.Review	
	filter := bson.M{}

	if req.BookingId != "" {
		filter["bookinng_id"] = req.BookingId
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

	cursor, err := r.collec.Find(context.TODO(), filter,
	options.Find().SetSkip(offset).SetLimit(limit))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var review pb.Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, &review)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &pb.ListAllReviewsResp{Reviews: reviews}, nil	
}