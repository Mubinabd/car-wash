package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type ProviderService struct {
	pb.UnimplementedProviderServiceServer
	Repo dbstore.Storage
}

func NewProviderService(repo dbstore.Storage) *ProviderService {
	return &ProviderService{
		Repo: repo,
	}
}

func (s *ProviderService) RegisterProvider(ctx context.Context, req *pb.RegisterProviderReq) (*pb.Empty, error) {
	return s.Repo.Provider().RegisterProvider(req)
}

func (s *ProviderService) GetProvider(ctx context.Context, req *pb.GetById) (*pb.Provider, error) {
	return s.Repo.Provider().GetProvider(req)
}

func (s *ProviderService) ListAllProviders(ctx context.Context, req *pb.ListAllProvidersReq) (*pb.ListAllProvidersResp, error) {
	return s.Repo.Provider().ListAllProviders(req)
}

func (s *ProviderService) UpdateProvider(ctx context.Context, req *pb.UpdateProviderReq) (*pb.UpdateProviderResp, error) {
	return s.Repo.Provider().UpdateProvider(req)
}

func (s *ProviderService) DeleteProvider(ctx context.Context, req *pb.DeleteProviderReq) (*pb.DeleteProviderResp, error) {
	return s.Repo.Provider().DeleteProvider(req)
}

func (s *ProviderService) SearchProviders(ctx context.Context, req *pb.SearchProvidersReq) (*pb.SearchProvidersResp, error) {
	return s.Repo.Provider().SearchProviders(req)
}