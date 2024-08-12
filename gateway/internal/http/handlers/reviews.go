package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
)
// @Router        /api/v1/review/add [POST]
// @Summary       CREATE review
// @Description   This API creates a review
// @Tags          carwash/review
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.AddReviewReq true "review"
// @Success       200 {object} string "message": "created successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) AddReview(c *gin.Context) {
	req := pb.AddReviewReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.Reviews.AddReview(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Create Reviews: Reviews created successfully: ")

	c.JSON(200, gin.H{"message": "Reviews created successfully"})

}
// Get review godoc
// @Summary      Get review
// @Description  This API Gets a  review
// @Tags         carwash/review
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "review Id"
// @Success      200 {object} pb.Review
// @Failure      400 {object} string "error"
// @Router       /api/v1/review/{id} [get]
func (h *Handlers) GetReview(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Reviews.GetReview(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Get Reviews: Reviews retrieved successfully: ", logrus.Fields{
		"Reviews_id": res.Id,
	})

	c.JSON(200, res)
}
// List all review godoc
// @Summary      List all review
// @Description  This API Lists a new review
// @Tags         carwash/review
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        booking_id query string false "Booking iD"
// @Param        provider_id query string false "provider ID"
// @Param        user_id query string false "User ID"
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200 {object} pb.ListAllReviewsResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/review [get]
func (h *Handlers) ListAllReviews(c *gin.Context) {
	var req pb.ListAllReviewsReq
	booking_id := c.Query("booking_id")
	provider_id := c.Query("_id")
	user_id := c.Query("user_id")
	limit := c.Query("limit")
	offset := c.Query("offset")

	req.BookingId = booking_id
	req.ProviderId = provider_id
	req.UserId = user_id

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
	res, err := h.Clients.Reviews.ListAllReviews(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("ListAllReviewss: Reviews retrieved successfully")
	c.JSON(200, res)
}
// Put review godoc
// @Summary      Put  review
// @Description  This API Put s a new review
// @Tags         carwash/review
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateReviewsReq reviewRequest true " review Request"
// @Success      200 {object} string "review updated successfully"
// @Failure      400 {object} string "error"
// @Router       /api/v1/review/{id} [put]

func (h *Handlers) UpdateReview(c *gin.Context) {
	var req pb.UpdateReviewsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Clients.Reviews.UpdateReview(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("update Reviews: Reviews retrieved successfully")

	c.JSON(200, gin.H{
		"message": "updated successfully",
	})
}
// Delete review godoc
// @Summary      Delete review
// @Description  This API deleted a new review
// @Tags         carwash/review
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "review Id"
// @Success      200 {object} string "review deleted successfully"
// @Failure      400 {object} string "error"
// @Router       /api/v1/review/{id} [delete]
func (h *Handlers) DeleteReview(c *gin.Context) {
	var req pb.DeleteReviewReq
	id := c.Param("id")
	req.Id = id

	_, err := h.Clients.Reviews.DeleteReview(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("delete Reviews: Reviews retrieved successfully")

	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})

}