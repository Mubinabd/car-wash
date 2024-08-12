package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
	Repo dbstore.Storage
}

func NewPaymentService(repo dbstore.Storage) *PaymentService {
	return &PaymentService{
		Repo: repo,
	}
}

func (s *PaymentService) AddPayment(ctx context.Context, req *pb.AddPaymentReq) (*pb.Empty, error) {
	return s.Repo.Payment().AddPayment(req)
}

func (s *PaymentService) GetPayment(ctx context.Context, req *pb.GetById) (*pb.GetPaymentResp, error) {
	return s.Repo.Payment().GetPayment(req)
}

func (s *PaymentService) ListAllPayments(ctx context.Context, req *pb.ListAllPaymentsReq) (*pb.ListAllPaymentsResp, error) {
	return s.Repo.Payment().ListAllPayments(req)
}
