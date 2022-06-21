package customers

import (
	"fmt"
	"net/http"

	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) Deletecustomer(c *gin.Context) {
	id := c.Param("id")

	var customer models.Customer

	if result := h.DB.First(&customer, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while deleting customer data", "error": result.Error})
		return

	}

	result := h.DB.Delete(&customer)
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
	return
}
