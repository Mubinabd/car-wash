package service

import (
	"context"
	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/dbstore"
)

type NotificationService struct {
	pb.UnimplementedNotificationServiceServer
	Repo dbstore.Storage
}

func NewNotificationService(repo dbstore.Storage) *NotificationService {
	return &NotificationService{
		Repo: repo,
	}
}

func (s *NotificationService) AddNotification(ctx context.Context, req *pb.AddNotificationReq) (*pb.Empty, error) {
	return s.Repo.Notification().AddNotification(req)
}

func (s *NotificationService) GetNotifications(ctx context.Context, req *pb.GetNotificationsReq) (*pb.Notification, error) {
	return s.Repo.Notification().GetNotifications(req)
}

func (s *NotificationService) MarkNotificationAsRead(ctx context.Context, req *pb.MarkNotificationAsReadReq) (*pb.MarkNotificationAsReadResp, error) {
	return s.Repo.Notification().MarkNotificationAsRead(req)
}
