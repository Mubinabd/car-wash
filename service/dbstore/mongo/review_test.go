package mongo

import (
	"context"
	"testing"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var testClient *mongo.Client
var testDatabase *mongo.Database
var reviewsManager *ReviewsManager

func setup() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	testClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	testDatabase = testClient.Database("testdb")
	reviewsManager = NewReviewsManager(testDatabase)

	testDatabase.Collection("reviews").Drop(context.TODO())
}

func teardown() {
	testClient.Disconnect(context.TODO())
}

func TestAddReview(t *testing.T) {
	setup()
	defer teardown()

	req := &pb.AddReviewReq{}
	_, err := reviewsManager.AddReview(req)
	assert.NoError(t, err)

	var count int64
	count, err = testDatabase.Collection("reviews").CountDocuments(context.TODO(), bson.M{})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}

func TestGetReview(t *testing.T) {
	setup()
	defer teardown()

	insertedID, err := testDatabase.Collection("reviews").InsertOne(context.TODO(), &pb.Review{})
	assert.NoError(t, err)

	oid := insertedID.InsertedID.(primitive.ObjectID)
	req := &pb.GetById{Id: oid.Hex()}
	review, err := reviewsManager.GetReview(req)
	assert.NoError(t, err)
	assert.NotNil(t, review)
}

func TestDeleteReview(t *testing.T) {
	setup()
	defer teardown()

	insertedID, err := testDatabase.Collection("reviews").InsertOne(context.TODO(), &pb.Review{})
	assert.NoError(t, err)

	oid := insertedID.InsertedID.(primitive.ObjectID)
	req := &pb.DeleteReviewReq{
		Id: oid.Hex(),
	}
	_, err = reviewsManager.DeleteReview(req)
	assert.NoError(t, err)
}

