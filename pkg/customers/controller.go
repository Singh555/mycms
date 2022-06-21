package customers

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/customer")
	routes.POST("/add", h.AddCustomer)
	routes.GET("/", h.GetCustomers)
	routes.GET("/:id", h.GetCustomer)

	routes.PUT("/update/:id", h.UpdateCustomer)
	routes.DELETE("/delete/:id", h.Deletecustomer)
}
