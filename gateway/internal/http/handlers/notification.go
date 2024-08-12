package handlers

import (
	"context"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router        /api/v1/notification/add [POST]
// @Summary       CREATE notification
// @Description   This API creates a notification
// @Tags          carwash/Notification
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.AddNotificationReq true "Notification"
// @Success       200 {object} string "message": "created successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) AddNotification(c *gin.Context) {
	var req pb.AddNotificationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// input, err := json.Marshal(req)
	// err = h.Clients.KafkaProducer.ProduceMessages("create-notification", input)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	log.Println("cannot produce messages via kafka", err)
	// 	return
	// }

	_, err := h.Clients.Notification.AddNotification(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Create Notification: notification created successfully: ")

	c.JSON(200, gin.H{
		"message": "created successfully",
	})
}

// @Router        /api/v1/notification/{id} [GET]
// @Summary       GET notification
// @Description   This API retrieves a notification by ID
// @Tags          carwash/Notification
// @Accept        json
// @Produce       json
// @Param         id path string true "Notification ID"
// @Success       200 {object} pb.GetNotificationsResp
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) GetNotifications(c *gin.Context) {
	var req pb.GetNotificationsReq
	id := c.Param("id")
	req.UserId = id
	res, err := h.Clients.Notification.GetNotifications(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Get Notification: notification retrieved successfully: ", logrus.Fields{
		"Notification_id": res.Notifications[0].Id,
	})
	c.JSON(200, res)
}

// @Router        /api/v1/notification/{id}/read [PUT]
// @Summary       MARK notification as read
// @Description   This API marks a notification as read
// @Tags          carwash/Notification
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.MarkNotificationAsReadReq true "Notification"
// @Success       200 {object} string "message": "updated successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) MarkNotificationAsRead(c *gin.Context) {
	var req pb.MarkNotificationAsReadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err := h.Clients.Notification.MarkNotificationAsRead(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Update Notification: notification updated successfully: ")

	c.JSON(200, gin.H{
		"message": "updated successfully",
	})
}
