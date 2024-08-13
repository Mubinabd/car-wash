package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type BookingsService struct {
	repo dbstore.Storage
	pb.UnimplementedBookingsServer
}

func NewBookingService(repo dbstore.Storage) *BookingsService {
	return &BookingsService{
		repo: repo,
	}
}

func (s *BookingsService) AddBooking(ctx context.Context, req *pb.AddBookingReq) (*pb.Empty, error) {
	return s.repo.Booking().AddBooking(req)
}

func (s *BookingsService) GetBooking(ctx context.Context, req *pb.GetById) (*pb.GetBookingResp, error) {
	return s.repo.Booking().GetBooking(req)
}

func (s *BookingsService) ListAllBookings(ctx context.Context, req *pb.ListAllBookingsReq) (*pb.ListAllBookingsResp, error) {
	return s.repo.Booking().ListAllBookings(req)
}

func (s *BookingsService) UpdateBooking(ctx context.Context, req *pb.UpdateBookingReq) (*pb.UpdateBookingResp, error) {
	return s.repo.Booking().UpdateBooking(req)
}

func (s *BookingsService) DeleteBooking(ctx context.Context, req *pb.DeleteBookingReq) (*pb.DeleteBookingResp, error) {
	return s.repo.Booking().DeleteBooking(req)
}
