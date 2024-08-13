package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	repo dbstore.Storage
}

func NewCartService(repo dbstore.Storage) *CartService {
	return &CartService{
		repo: repo,
	}
}

func (s *CartService) CreateCart(ctx context.Context, req *pb.CreateCartReq) (*pb.Empty, error) {
	return s.repo.Cart().CreateCart(req)
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetById) (*pb.Cart, error) {
	return s.repo.Cart().GetCart(req)
}
