package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type BookingsService struct {
	pb.UnimplementedBookingsServer
	Repo dbstore.Storage
}

func NewBookingService(repo dbstore.Storage) *BookingsService {
	return &BookingsService{
		Repo: repo,
	}
}

func (s *BookingsService) AddBooking(ctx context.Context, req *pb.AddBookingReq) (*pb.Empty, error) {
	return s.Repo.Booking().AddBooking(req)
}

func (s *BookingsService) GetBooking(ctx context.Context, req *pb.GetById) (*pb.Booking, error) {
	return s.Repo.Booking().GetBooking(req)
}

func (s *BookingsService) ListAllBookings(ctx context.Context, req *pb.ListAllBookingsReq) (*pb.ListAllBookingsResp, error) {
	return s.Repo.Booking().ListAllBookings(req)
}

func (s *BookingsService) UpdateBooking(ctx context.Context, req *pb.UpdateBookingReq) (*pb.UpdateBookingResp, error) {
	return s.Repo.Booking().UpdateBooking(req)
}

func (s *BookingsService) DeleteBooking(ctx context.Context, req *pb.DeleteBookingReq) (*pb.DeleteBookingResp, error) {
	return s.Repo.Booking().DeleteBooking(req)
}

func (s *BookingsService) GetBookingsByProvider(ctx context.Context, req *pb.BookingsByProviderReq) (*pb.BookingsByProviderResp, error) {
	return s.Repo.Booking().GetBookingsByProvider(req)
}