package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
)

// @Router        /api/v1/service/add [POST]
// @Summary       CREATE service
// @Description   This API creates a service
// @Tags          carwash/service
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.AddServiceReq true "service"
// @Success       200 {object} string "message": "created successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) AddService(c *gin.Context) {
	req := pb.AddServiceReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.Service.AddService(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Create Service: Service created successfully: ")

	c.JSON(200, gin.H{"message": "Service created successfully"})

}
// Get service godoc
// @Summary      Get service
// @Description  This API Gets a  service
// @Tags         carwash/service
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "service Id"
// @Success      200 {object} pb.GetServicesResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/service/{id} [get]
func (h *Handlers) GetService(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Service.GetServices(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Get Service: Service retrieved successfully: ", logrus.Fields{
		"Service_id": res.Services.Id,
	})

	c.JSON(200, res)
}
// List all service godoc
// @Summary      List all service
// @Description  This API Lists a new service
// @Tags         carwash/service
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        name query string false "Name"
// @Param        description query string false "Description"
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200 {object} pb.ListAllServicesResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/service [get]
func (h *Handlers) ListAllServices(c *gin.Context) {
	var req pb.ListAllServicesReq
	name := c.Query("name")
	description := c.Query("description")
	limit := c.Query("limit")
	offset := c.Query("offset")

	req.Name = name
	req.Description = description

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
	res, err := h.Clients.Service.ListAllServices(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("ListAllServices: Service retrieved successfully")
	c.JSON(200, res)
}
// Put service godoc
// @Summary      Put  service
// @Description  This API Put s a new service
// @Tags         carwash/service
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateServiceReq true "Put  service Request"
// @Success      200 {object} string "service updated successfully"
// @Failure      400 {object} string "error"
// @Router       /api/v1/service/{id} [put]
func (h *Handlers) UpdateService(c *gin.Context) {
	var req pb.UpdateServiceReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.Clients.Service.UpdateService(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("update Service: Service retrieved successfully")

	c.JSON(200, gin.H{
		"message": "updated successfully",
	})
}
// Delete service godoc
// @Summary      Delete service
// @Description  This API deleted a new service
// @Tags         carwash/service
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "service Id"
// @Success      200 {object} string "service deleted successfully"
// @Failure      400 {object} string "error"
// @Router       /api/v1/service/{id} [delete]
func (h *Handlers) DeleteService(c *gin.Context) {
	var req pb.DeleteServiesReq
	id := c.Param("id")
	req.Id = id

	_, err := h.Clients.Service.DeleteService(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("delete Service: Service retrieved successfully")

	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})

}
// Get service id by service godoc
// @Summary      Get service
// @Description  This API gets a  service
// @Tags         carwash/service
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        name path string true "name"
// @Param        description path string true "description"
// @Success      200 {object} pb.SearchServicessResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/service/search [get]
func (h *Handlers) SearchServices(c *gin.Context) {
	var req pb.SearchServicessReq
	name := c.Query("name")
	description := c.Query("description")

	req.Name = name
	req.Description = description

	res, err := h.Clients.Service.SearchServices(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}
// Get service id by service godoc
// @Summary      Get service
// @Description  This API gets a  service
// @Tags         carwash/service
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        min_price path int true "min_price"
// @Param        max_price path int true "max_price"
// @Success      200 {object} pb.GetServicesByPriceRangeResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/service/priceRange [get]
func (h *Handlers) GetServicesByPriceRange(c *gin.Context) {
	var req pb.GetServicesByPriceRangeReq
	minpriceSTR := c.Query("min_price")
	minprice,err := strconv.Atoi(minpriceSTR)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	maxpriceSTR := c.Query("max_price")
	maxprice,err := strconv.Atoi(maxpriceSTR)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	req.MaxPrice = float32(maxprice)
	req.MinPrice = float32(minprice)
	res, err := h.Clients.Service.GetServicesByPriceRange(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}