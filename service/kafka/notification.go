package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/service"
)

func NotificationHandler(notifservice *service.NotificationService) func(message []byte) {
	return func(message []byte) {
		var notif pb.AddNotificationReq
		if err := json.Unmarshal(message, &notif); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		respnotif, err := notifservice.AddNotification(context.Background(), &notif)
		if err != nil {
			log.Printf("Cannot create notif via Kafka: %v", err)
			return
		}
		log.Printf("Created notif: %+v",respnotif)
	}
}
