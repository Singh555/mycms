package customers

import (
	"net/http"

	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetCustomers(c *gin.Context) {
	var customers []models.Customer

	if result := h.DB.Limit(10).Select("id", "first_name", "last_name", "email", "mobile", "address", "status", "created_at", "updated_at").Order("id DESC").Find(&customers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while getting customer data", "error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer data found", "data": &customers})
}
