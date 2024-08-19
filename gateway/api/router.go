package http

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Mubinabd/car-wash/api/handlers"
	middlerware "github.com/Mubinabd/car-wash/api/middleware"
	_ "github.com/Mubinabd/car-wash/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @BasePath  /api/v1
// @description                 Description for what is this security definition being used
func NewRouter(h *handlers.Handlers) *gin.Engine {
	router := gin.Default()

	enforcer, err := casbin.NewEnforcer("./load/model.conf", "./load/policy.csv")

	if err != nil {
		log.Fatal(err)
	}
	sw := router.Group("/")
	sw.Use(middlerware.NewAuth(enforcer))

	// CORS configuration
	// corsConfig := cors.Config{
	// 	AllowOrigins:     []string{"http://localhost", "http://localhost:8090"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	// 	AllowCredentials: true,
	// }
	// router.Use(cors.New(corsConfig))

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Cart routes
		cart := v1.Group("/cart")
		{
			cart.POST("/add", h.CreateCart)
			cart.GET("/:id", h.GetCart)
		}

		// Services routes
		services := v1.Group("/service")
		{
			services.POST("/add", h.AddService)
			services.GET("/:id", h.GetServices)
			services.GET("", h.ListAllServices)
			services.PUT("/:id", h.UpdateService)
			services.GET("/search", h.SearchServices)
			services.GET("/priceRange", h.GetServicesByPriceRange)
		}

		// Reviews routes
		reviews := v1.Group("/review")
		{
			reviews.POST("/add", h.AddReview)
			reviews.GET("/:id", h.GetReview)
			reviews.GET("", h.ListAllReviews)
			reviews.PUT("/:id", h.UpdateReview)
			reviews.DELETE("/:id", h.DeleteReview)
		}

		// Provider routes
		provider := v1.Group("/provider")
		{
			provider.POST("/add", h.RegisterProvider)
			provider.GET("/:id", h.GetProvider)
			provider.GET("", h.ListAllProviders)
			provider.PUT("/:id", h.UpdateProvider)
			provider.DELETE("/:id", h.DeleteProvider)
			provider.GET("/search", h.SearchProviders)
		}

		// Booking routes
		booking := v1.Group("/booking")
		{
			booking.POST("/add", h.AddBooking)
			booking.GET("/:id", h.GetBooking)
			booking.GET("", h.ListAllBookings)
			booking.PUT("/:id", h.UpdateBooking)
			booking.DELETE("/:id", h.DeleteBooking)
		}

		// Payment routes
		payment := v1.Group("/payment")
		{
			payment.POST("/add", h.AddPayment)
			payment.GET("/:id", h.GetPayment)
			payment.GET("", h.ListAllPayments)
		}

	}

	return router
}
