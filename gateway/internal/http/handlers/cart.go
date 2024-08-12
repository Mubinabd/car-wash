package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	pb "github.com/Mubinabd/car-wash/genproto"
	"github.com/Mubinabd/car-wash/logger"
)
// @Router 				/api/v1/cart/add [POST]
// @Summary 			CREATE cart
// @Description 		This API creates a new cart.
// @Tags 				carwash/Cart
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 				data body pb.CreateCartReq true "Cart data"
// @Success 			200 {object} string "message": "created successfully"
// @Failure 			400 {object} string "error": "error message"
func (h *Handlers) CreateCart(c *gin.Context) {
	req := pb.CreateCartReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.Cart.CreateCart(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Create Cart: Cart created successfully: ")

	c.JSON(200, gin.H{"message": "Cart created successfully"})

}

// @Router 				/api/v1/cart/{id} [GET]
// @Summary 			GET cart by ID
// @Description 		Retrieve a specific cart by its ID.
// @Tags 				carwash/Cart
// @Accept 				json
// @Produce 			json
// @Param 				id path string true "Cart ID"
// @Success 			200 {object} pb.Cart
// @Failure 			400 {object} string "error": "error message"
func (h *Handlers) GetCart(c *gin.Context) {
	req := pb.GetById{}
	id := c.Param("id")

	req.Id = id

	res, err := h.Clients.Cart.GetCart(context.Background(), &req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Get Cart: Cart retrieved successfully: ", logrus.Fields{
		"Cart_id": res.Id,
	})

	c.JSON(200, res)
}
