package customers

import (
	"github.com/Singh555/mycms/common/middlewares"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

//creating a handler for gorm instance
type handler struct {
	DB *gorm.DB
}

//function to register customer module routes
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	api := r.Group("/api")
	{
		routes := api.Group("/auth")
		{
			routes.POST("/register", h.AddCustomer)
			routes.POST("/login", h.LoginCustomer)

		}
		//secured routes request must have a jwt token to access these endpoints

		secured := api.Group("/customer").Use(middlewares.Auth())
		{
			secured.GET("list/", h.GetCustomers)
			secured.GET("/:id", h.GetCustomer)
			secured.POST("profile/", h.GetProfile)
			secured.PUT("/update/", h.UpdateCustomer)
			secured.DELETE("/delete/:id", h.Deletecustomer)
		}

	}

}
