package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
)

// @Router        /api/v1/provider/add [POST]
// @Summary       CREATE provider
// @Description   This API creates a provider
// @Tags          carwash/provider
// @Accept        json
// @Produce       json
// @Security      BearerAuth
// @Param         data body pb.RegisterProviderReq true "Provider"
// @Success       200 {object} string "message": "created successfully"
// @Failure       400 {object} string "error": "error message"
func (h *Handlers) RegisterProvider(c *gin.Context) {
	var req pb.RegisterProviderReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if h.Clients.ProviderClient == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ProviderClient is not initialized"})
		return
	}
	input, err := json.Marshal(req)
	err = h.Clients.KafkaProducer.ProduceMessages("cr-provider", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}
	// _, err := h.Clients.ProviderClient.RegisterProvider(context.Background(), &req)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"message": "Provider registered successfully"})
}

// Get provider godoc
// @Summary      Get provider
// @Description  This API Gets a  provider
// @Tags         carwash/provider
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "provider Id"
// @Success      200 {object} pb.GetProviderResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/provider/{id} [get]
func (h *Handlers) GetProvider(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.ProviderClient.GetProvider(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Get Provider: provider retrieved successfully", logrus.Fields{
		"provider_id": res.Provider.Id,
	})

	c.JSON(200, res)
}

// List all provider godoc
// @Summary      List all provider
// @Description  This API Lists a new provider
// @Tags         carwash/provider
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        company_name query string false "Company Name"
// @Param        description query string false "Description"
// @Param        limit query int false "Limit"
// @Param        offset query int false "Offset"
// @Success      200 {object} pb.ListAllProvidersResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/provider [get]
func (h *Handlers) ListAllProviders(c *gin.Context) {
	var req pb.ListAllProvidersReq
	company_name := c.Query("compnay_name")
	description := c.Query("description")
	limit := c.Query("limit")
	offset := c.Query("offset")

	req.CompanyName = company_name
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
	res, err := h.Clients.ProviderClient.ListAllProviders(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Info("ListAllProviders: provider retrieved successfully")
	c.JSON(200, res)
}

// Put provider godoc
// @Summary      Put  provider
// @Description  This API Put s a new provider
// @Tags         carwash/provider
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateProviderReq true "Put  provider Request"
// @Success      200 {object} string "provider updated successfully"
// @Failure      400 {object} string "error"
// @Router       /api/v1/provider/{id} [put]
func (h *Handlers) UpdateProvider(c *gin.Context) {
	var req pb.UpdateProviderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	input, err := json.Marshal(req)
	err = h.Clients.KafkaProducer.ProduceMessages("up-provider", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	// _, err := h.Clients.ProviderClient.UpdateProvider(context.Background(), &req)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	logger.Info("update provider: provider retrieved successfully")

	c.JSON(200, gin.H{
		"message": "updated successfully",
	})
}

// Delete provider godoc
// @Summary      Delete provider
// @Description  This API deleted a new provider
// @Tags         carwash/provider
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "provider Id"
// @Success      200 {object} string "provider deleted successfully"
// @Failure      400 {object} string "error"
// @Router       /api/v1/provider/{id} [delete]
func (h *Handlers) DeleteProvider(c *gin.Context) {
	var req pb.DeleteProviderReq
	id := c.Param("id")
	req.Id = id

	input, err := json.Marshal(req)
	err = h.Clients.KafkaProducer.ProduceMessages("dl-provider", input)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		log.Println("cannot produce messages via kafka", err)
		return
	}

	// _, err := h.Clients.ProviderClient.DeleteProvider(context.Background(), &req)
	// if err != nil {
	// 	c.JSON(400, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	logger.Info("delete provider: provider retrieved successfully")

	c.JSON(200, gin.H{
		"message": "deleted successfully",
	})

}

// Get Provider id by provider godoc
// @Summary      Get provider
// @Description  This API gets a  provider
// @Tags         carwash/provider
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        name path string true "name"
// @Param        description path string true "description"
// @Param        user_id path string true "user_id"
// @Success      200 {object} pb.SearchProvidersResp
// @Failure      400 {object} string "error"
// @Router       /api/v1/provider/search [get]
func (h *Handlers) SearchProviders(c *gin.Context) {
	var req pb.SearchProvidersReq
	name := c.Query("name")
	description := c.Query("description")
	userid := c.Query("user_id")

	req.CompanyName = name
	req.Description = description
	req.UserId = userid

	res, err := h.Clients.ProviderClient.SearchProviders(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}
