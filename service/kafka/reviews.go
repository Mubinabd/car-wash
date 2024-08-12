package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/service"
)

func ReviewHandler(reviewservice *service.ReviewService) func(message []byte) {
	return func(message []byte) {
		var review pb.AddReviewReq
		if err := json.Unmarshal(message, &review); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respreview, err := reviewservice.AddReview(context.Background(), &review)
		if err != nil {
			log.Printf("Cannot create review via Kafka: %v", err)
			return
		}
		log.Printf("Created review: %+v",respreview)
	}
}

func UpdatereviewHandler(reviewservice *service.ReviewService) func(message []byte) {
	return func(message []byte) {
		var review pb.UpdateReviewsReq
		if err := json.Unmarshal(message, &review); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respreview, err := reviewservice.UpdateReview(context.Background(), &review)
		if err != nil {
			log.Printf("Cannot create review via Kafka: %v", err)
			return
		}
		log.Printf("Created review: %+v",respreview)
	}
}

func DeletereviewHandler(reviewservice *service.ReviewService) func(message []byte) {
	return func(message []byte) {
		var review pb.DeleteReviewReq
		if err := json.Unmarshal(message, &review); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respreview, err := reviewservice.DeleteReview(context.Background(), &review)
		if err != nil {
			log.Printf("Cannot create review via Kafka: %v", err)
			return
		}
		log.Printf("Created review: %+v",respreview)
	}
}
