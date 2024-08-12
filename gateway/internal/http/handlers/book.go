package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
)

// AddBooking godoc
// @Summary      Add booking
// @Description  This API adds a new booking
// @Tags         carwash/booking
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.AddBookingReq true "Add booking Request"
// @Success      200 {object} string "message": "booking created successfully"
// @Failure      400 {object} string "error": "error description"
// @Failure      500 {object} string "error": "error description"
// @Router       /api/v1/booking/add [post]
func (h *Handlers) AddBooking(c *gin.Context) {
	var req pb.AddBookingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.BookingClient.AddBooking(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Create Booking: booking created successfully")

	c.JSON(200, gin.H{"message": "booking created successfully"})
}

// GetBooking godoc
// @Summary      Get booking
// @Description  This API retrieves a booking by its ID
// @Tags         carwash/booking
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Booking Id"
// @Success      200 {object} pb.GetBookingResp
// @Failure      400 {object} string "error": "error description"
// @Failure      500 {object} string "error": "error description"
// @Router       /api/v1/booking/{id} [get]
func (h *Handlers) GetBooking(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.BookingClient.GetBooking(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Get booking: Booking retrieved successfully", logrus.Fields{
		"Booking_id": res.Booking.Id,
	})

	c.JSON(200, res)
}

// ListAllBookings godoc
// @Summary      List all bookings
// @Description  This API lists all bookings with optional filters
// @Tags         carwash/booking
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        user_id query string false "User ID"
// @Param        status query string false "Status"
// @Param        provider_id query string false "Provider ID"
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200 {object} pb.ListAllBookingsResp
// @Failure      400 {object} string "error": "error description"
// @Router       /api/v1/booking [get]
func (h *Handlers) ListAllBookings(c *gin.Context) {
	var req pb.ListAllBookingsReq
	status := c.Query("status")
	user_id := c.Query("user_id")
	provider_id := c.Query("provider_id")
	limit := c.Query("limit")
	offset := c.Query("offset")

	req.Status = status
	req.UserId = user_id
	req.ProviderId = provider_id

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

	res, err := h.Clients.BookingClient.ListAllBookings(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logger.Info("ListAllBookings: Bookings retrieved successfully")
	c.JSON(200, res)
}

// UpdateBooking godoc
// @Summary      Update booking
// @Description  This API updates an existing booking
// @Tags         carwash/booking
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Booking Id"
// @Param        request body pb.UpdateBookingReq true "Update booking Request"
// @Success      200 {object} string "message": "booking updated successfully"
// @Failure      400 {object} string "error": "error description"
// @Failure      500 {object} string "error": "error description"
// @Router       /api/v1/booking/{id} [put]
func (h *Handlers) UpdateBooking(c *gin.Context) {
	var req pb.UpdateBookingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	req.Id = id

	_, err := h.Clients.BookingClient.UpdateBooking(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Update Booking: Booking updated successfully")

	c.JSON(200, gin.H{"message": "booking updated successfully"})
}

// DeleteBooking godoc
// @Summary      Delete booking
// @Description  This API deletes a booking by its ID
// @Tags         carwash/booking
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Booking Id"
// @Success      200 {object} string "message": "booking deleted successfully"
// @Failure      400 {object} string "error": "error description"
// @Failure      500 {object} string "error": "error description"
// @Router       /api/v1/booking/{id} [delete]
func (h *Handlers) DeleteBooking(c *gin.Context) {
	var req pb.DeleteBookingReq
	id := c.Param("id")
	req.Id = id

	_, err := h.Clients.BookingClient.DeleteBooking(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Delete Booking: Booking deleted successfully")

	c.JSON(200, gin.H{"message": "booking deleted successfully"})
}

// GetBookingsByProvider godoc
// @Summary      Get bookings by provider
// @Description  This API retrieves bookings for a specific provider
// @Tags         carwash/booking
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        provider_id query string true "Provider ID"
// @Success      200 {object} pb.BookingsByProviderResp
// @Failure      400 {object} string "error": "error description"
// @Failure      500 {object} string "error": "error description"
// @Router       /api/v1/booking/provider [get]
func (h *Handlers) GetBookingsByProvider(c *gin.Context) {
	var req pb.BookingsByProviderReq
	provider_id := c.Query("provider_id")

	req.ProviderId = provider_id

	res, err := h.Clients.BookingClient.GetBookingsByProvider(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}