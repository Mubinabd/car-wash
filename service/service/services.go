package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type ServiceService struct {
	pb.UnimplementedServicesServiceServer
	repo dbstore.Storage
}

func NewServiceService(repo dbstore.Storage) *ServiceService {
	return &ServiceService{
		repo: repo,
	}
}

func (s *ServiceService) AddService(ctx context.Context, req *pb.AddServiceReq) (*pb.Empty, error) {
	return s.repo.Servicee().AddService(req)
}

func (s *ServiceService) GetServices(ctx context.Context, req *pb.GetById) (*pb.GetServicesResp, error) {
	return s.repo.Servicee().GetServices(req)
}

func (s *ServiceService) ListAllServices(ctx context.Context, req *pb.ListAllServicesReq) (*pb.ListAllServicesResp, error) {
	return s.repo.Servicee().ListAllServices(req)
}

func (s *ServiceService) UpdateService(ctx context.Context, req *pb.UpdateServiceReq) (*pb.UpdateServiceResp, error) {
	return s.repo.Servicee().UpdateService(req)
}

func (s *ServiceService) DeleteService(ctx context.Context, req *pb.DeleteServiesReq) (*pb.DeleteServiesResp, error) {
	return s.repo.Servicee().DeleteService(req)
}

func (s *ServiceService) SearchServices(ctx context.Context, req *pb.SearchServicessReq) (*pb.SearchServicessResp, error) {
	return s.repo.Servicee().SearchServices(req)
}
func (s *ServiceService) GetServicesByPriceRange(ctx context.Context, req *pb.GetServicesByPriceRangeReq) (*pb.GetServicesByPriceRangeResp, error) {
	return s.repo.Servicee().GetServicesByPriceRange(req)
}