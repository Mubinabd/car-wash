package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type NotificationService struct {
	pb.UnimplementedNotificationServiceServer
	repo dbstore.Storage
}

func NewNotificationService(repo dbstore.Storage) *NotificationService {
	return &NotificationService{
		repo: repo,
	}
}

func (s *NotificationService) AddNotification(ctx context.Context, req *pb.AddNotificationReq) (*pb.Empty, error) {
	return s.repo.Notification().AddNotification(req)
}

func (s *NotificationService) GetNotifications(ctx context.Context, req *pb.GetNotificationsReq) (*pb.GetNotificationsResp, error) {
	return s.repo.Notification().GetNotifications(req)
}

func (s *NotificationService) MarkNotificationAsRead(ctx context.Context, req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error) {
	return s.repo.Notification().MarkNotificationAsRead(req)
}
