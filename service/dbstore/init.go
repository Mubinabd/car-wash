package dbstore

import (
	pb "github.com/Mubinabd/car-wash/genproto"
)
type Storage interface {
	Booking() BookingI
	Cart() CartI
	Notification() NotificationI
	Payment() PaymentI
	Provider() ProviderI
	Review() ReviewI
	Service() ServiceI
}	

type BookingI interface {
	AddBooking(req *pb.AddBookingReq) (*pb.Empty, error)
	GetBooking(req *pb.GetById) (*pb.Booking, error)
	ListAllBookings(req *pb.ListAllBookingsReq) (*pb.ListAllBookingsResp, error)
	UpdateBooking(req *pb.UpdateBookingReq) (*pb.UpdateBookingResp, error)
	DeleteBooking(req *pb.DeleteBookingReq) (*pb.DeleteBookingResp, error)
	GetBookingsByProvider(req *pb.BookingsByProviderReq) (*pb.BookingsByProviderResp, error)
}

type CartI interface {
	CreateCart(req *pb.CreateCartReq) (*pb.Empty, error)
	GetCart(req *pb.GetById) (*pb.Cart, error)
}

type NotificationI interface {
	AddNotification(req *pb.AddNotificationReq) (*pb.Empty, error)
	GetNotifications(req *pb.GetNotificationsReq) (*pb.Notification, error)
	MarkNotificationAsRead(req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error)
}

type PaymentI interface {
	AddPayment(req *pb.AddPaymentReq) (*pb.Empty, error)
	GetPayment(req *pb.GetById) (*pb.GetPaymentResp, error)
	ListAllPayments(req *pb.ListAllPaymentsReq) (*pb.ListAllPaymentsResp, error)
}

type ProviderI interface {
	RegisterProvider(req *pb.RegisterProviderReq) (*pb.Empty, error)
	GetProvider(req *pb.GetById) (*pb.Provider, error)
	ListAllProviders(req *pb.ListAllProvidersReq) (*pb.ListAllProvidersResp, error)
	UpdateProvider(req *pb.UpdateProviderReq) (*pb.UpdateProviderResp, error)
	DeleteProvider(req *pb.DeleteProviderReq) (*pb.DeleteProviderResp, error)
	SearchProviders(req *pb.SearchProvidersReq) (*pb.SearchProvidersResp, error)
}

type ReviewI interface {
	AddReview(req *pb.AddReviewReq) (*pb.Empty, error)
	GetReview(req *pb.GetById) (*pb.Review, error)
	UpdateReview(req *pb.UpdateReviewsReq) (*pb.UpdateReviewsResp, error)
	DeleteReview(req *pb.DeleteReviewReq) (*pb.DeleteReviewResp, error)
	ListAllReviews(req *pb.ListAllReviewsReq) (*pb.ListAllReviewsResp, error)
}

type ServiceI interface {
	AddService(req *pb.AddServiceReq) (*pb.Empty, error)
	GetService(req *pb.GetById) (*pb.Services, error)
	ListAllServices(req *pb.ListAllServicesReq) (*pb.ListAllServicesResp, error)
	UpdateService(req *pb.UpdateServiceReq) (*pb.UpdateServiceResp, error)
	DeleteService(req *pb.DeleteServiesReq) (*pb.DeleteServiesResp, error)
	SearchServices(req *pb.SearchServicessReq) (*pb.SearchServicessResp, error)
	GetServicesByPriceRange(req *pb.GetServicesByPriceRangeReq) (*pb.GetServicesByPriceRangeResp, error)
}
