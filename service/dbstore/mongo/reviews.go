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
	var review pb.Review
	collection := r.collec.Database().Collection("reviews")
	filter := bson.M{"id": req.Id}

	err := collection.FindOne(context.TODO(), filter).Decode(&review)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no review found with the given ID")
		}
		return nil, err
	}
	logger.Info("Review retrieved successfully", logrus.Fields{
		"review_id": review.Id,
		"rating":    review.Rating,
		"comment":   review.Comment,
	})
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
	collection := r.collec.Database().Collection("reviews")	
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

	cursor, err := collection.Find(context.TODO(), filter,
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

	return &pb.ListAllReviewsResp{Reviews: reviews}, nil	
}