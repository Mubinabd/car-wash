package handlers

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	input, err := json.Marshal(req)
	err = h.Clients.KafkaProducer.ProduceMessages("cr-service", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	// _, err := h.Clients.Service.AddService(context.Background(), &req)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

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
func (h *Handlers) GetServices(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Service.GetServices(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200 {object} pb.ListAllServicesResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/service [get]
func (h *Handlers) ListAllServices(c *gin.Context) {
	var req pb.ListAllServicesReq
	limit := c.Query("limit")
	offset := c.Query("offset")

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

	redisKey := "my_services"

	cachedData, err := h.Clients.RedisClient.Get(context.Background(), redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			c.JSON(404, gin.H{"error": "Popular services not found in cache"})
			return
		}
		logger.Error("Error fetching popular services from Redis: ", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	var popularServices []*pb.Services
	err = json.Unmarshal([]byte(cachedData), &popularServices)
	if err != nil {
		logger.Error("Error unmarshalling popular services cached data: ", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	var filteredServices []pb.Services
	for _, service := range popularServices {
		if service.Price >= 50 {
			filteredServices = append(filteredServices, *service)
		}
	}

	c.JSON(200, filteredServices)
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

	input, err := json.Marshal(req)
	err = h.Clients.KafkaProducer.ProduceMessages("up-service", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	// _, err := h.Clients.Service.UpdateService(context.Background(), &req)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

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

	input, err := json.Marshal(req)
	err = h.Clients.KafkaProducer.ProduceMessages("dl-service", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	// _, err := h.Clients.Service.DeleteService(context.Background(), &req)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

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
// @Success      200 {object} pb.SearchServicessResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/service/search [get]
func (h *Handlers) SearchServices(c *gin.Context) {
	var req pb.SearchServicessReq
	name := c.Query("name")

	req.Name = name

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
	if minpriceSTR == "" {
		minpriceSTR = "0"
	}
	minprice, err := strconv.Atoi(minpriceSTR)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid min_price parameter: " + err.Error(),
		})
		return
	}

	maxpriceSTR := c.Query("max_price")
	if maxpriceSTR == "" {
		maxpriceSTR = "0"
	}
	maxprice, err := strconv.Atoi(maxpriceSTR)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid max_price parameter: " + err.Error(),
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
