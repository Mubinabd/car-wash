package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type CartService struct {
	pb.UnimplementedCartServiceServer
	Repo dbstore.Storage
}

func NewCartService(repo dbstore.Storage) *CartService {
	return &CartService{
		Repo: repo,
	}
}

func (s *CartService) CreateCart(ctx context.Context, req *pb.CreateCartReq) (*pb.Empty, error) {
	return s.Repo.Cart().CreateCart(req)
}

func (s *CartService) GetCart(ctx context.Context, req *pb.GetById) (*pb.Cart, error) {
	return s.Repo.Cart().GetCart(req)
}
