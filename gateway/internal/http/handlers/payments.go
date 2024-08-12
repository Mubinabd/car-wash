package handlers

import (
	"context"
	"strconv"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// @Router        /api/v1/payment/add [POST]
// @Summary       CREATE Payment
// @Description   This API creates a Payment
// @Tags          carwash/Payment
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.AddPaymentReq true "Payment"
// @Success       200 {object} string "message": "created successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) AddPayment(c *gin.Context) {
	var req pb.AddPaymentReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	// input, err := json.Marshal(req)
	// err = h.Clients.KafkaProducer.ProduceMessages("create-Payment", input)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	log.Println("cannot produce messages via kafka", err)
	// 	return
	// }

	_, err := h.Clients.Payments.AddPayment(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Create Payment: Payment created successfully: ")

	c.JSON(200, gin.H{
		"message": "created successfully",
	})
}

// @Router        /api/v1/payment/{id} [GET]
// @Summary       GET Payment
// @Description   This API retrieves a Payment by ID
// @Tags          carwash/Payment
// @Accept        json
// @Produce       json
// @Param         id path string true "Payment ID"
// @Success       200 {object} pb.Payment
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) GetPayment(c *gin.Context) {
	var req pb.GetById
	id := c.Param("id")
	req.Id = id
	res, err := h.Clients.Payments.GetPayment(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Get Payment: Payment retrieved successfully: ", logrus.Fields{
		"Payment_id": res.Payment.Id,
	})
	c.JSON(200, res)
}

// @Router        /api/v1/payment [GET]
// @Summary       GET all Payments
// @Description   This API retrieves all Payments with optional filters
// @Tags          carwash/Payment
// @Accept        json
// @Produce       json
// @Param         booking_id query string false "Booking ID"
// @Param         status query string false "Status"
// @Param         limit query int false "Limit"
// @Param         offset query int false "Offset"
// @Success       200 {object} pb.ListAllPaymentsResp
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) ListAllPayments(c *gin.Context) {
	var req pb.ListAllPaymentsReq
	booking_id := c.Query("booking_id")
	status := c.Query("status")
	limit := c.Query("limit")
	offset := c.Query("offset")

	req.BookingId = booking_id
	req.Status = status

	if limit != "" {
		limitValue, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid limit value"})
			return
		}
		if req.Filter == nil {
			req.Filter = &pb.Filter{}
		}
		req.Filter.Limit = int32(limitValue)
	}

	if offset != "" {
		offsetValue, err := strconv.Atoi(offset)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid offset value"})
			return
		}
		if req.Filter == nil {
			req.Filter = &pb.Filter{}
		}
		req.Filter.Offset = int32(offsetValue)
	}
	res, err := h.Clients.Payments.ListAllPayments(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info("Get Payment: Payment retrieved successfully: ", logrus.Fields{
		"Payment_id": res.Payments[0].Id,
	})
	
	c.JSON(200, res)
}
