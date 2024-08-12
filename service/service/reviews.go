package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type ReviewService struct {
	pb.UnimplementedReviewServiceServer
	Repo dbstore.Storage
}

func NewReviewService(repo dbstore.Storage) *ReviewService {
	return &ReviewService{
		Repo: repo,
	}
}

func (s *ReviewService) AddReview(ctx context.Context, req *pb.AddReviewReq) (*pb.Empty, error) {
	return s.Repo.Review().AddReview(req)
}

func (s *ReviewService) GetReview(ctx context.Context, req *pb.GetById) (*pb.Review, error) {
	return s.Repo.Review().GetReview(req)
}

func (s *ReviewService) ListAllReviews(ctx context.Context, req *pb.ListAllReviewsReq) (*pb.ListAllReviewsResp, error) {
	return s.Repo.Review().ListAllReviews(req)
}

func (s *ReviewService) UpdateReview(ctx context.Context, req *pb.UpdateReviewsReq) (*pb.UpdateReviewsResp, error) {
	return s.Repo.Review().UpdateReview(req)
}

func (s *ReviewService) DeleteReview(ctx context.Context, req *pb.DeleteReviewReq) (*pb.DeleteReviewResp, error) {
	return s.Repo.Review().DeleteReview(req)
}
