package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type ServiceService struct {
	pb.UnimplementedServicesServiceServer
	Repo dbstore.Storage
}

func NewServiceService(repo dbstore.Storage) *ServiceService {
	return &ServiceService{
		Repo: repo,
	}
}

func (s *ServiceService) AddService(ctx context.Context, req *pb.AddServiceReq) (*pb.Empty, error) {
	return s.Repo.Service().AddService(req)
}

func (s *ServiceService) GetService(ctx context.Context, req *pb.GetById) (*pb.Services, error) {
	return s.Repo.Service().GetService(req)
}

func (s *ServiceService) ListAllServices(ctx context.Context, req *pb.ListAllServicesReq) (*pb.ListAllServicesResp, error) {
	return s.Repo.Service().ListAllServices(req)
}

func (s *ServiceService) UpdateService(ctx context.Context, req *pb.UpdateServiceReq) (*pb.UpdateServiceResp, error) {
	return s.Repo.Service().UpdateService(req)
}

func (s *ServiceService) DeleteService(ctx context.Context, req *pb.DeleteServiesReq) (*pb.DeleteServiesResp, error) {
	return s.Repo.Service().DeleteService(req)
}

func (s *ServiceService) SearchServices(ctx context.Context, req *pb.SearchServicessReq) (*pb.SearchServicessResp, error) {
	return s.Repo.Service().SearchServices(req)
}
func (s *ServiceService) GetServicesByPriceRange(ctx context.Context, req *pb.GetServicesByPriceRangeReq) (*pb.GetServicesByPriceRangeResp, error) {
	return s.Repo.Service().GetServicesByPriceRange(req)
}