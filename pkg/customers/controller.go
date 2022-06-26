package customers

import (
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

	routes := r.Group("/customer")
	routes.POST("/add", h.AddCustomer)
	routes.GET("/", h.GetCustomers)
	routes.GET("/:id", h.GetCustomer)

	routes.PUT("/update/", h.UpdateCustomer)
	routes.DELETE("/delete/:id", h.Deletecustomer)
	routes.POST("/login", h.LoginCustomer)
}
