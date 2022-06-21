package customers

import (
	"net/http"

	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateCustomerRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	//Mobile    string `json:"mobile"`
	Address string `json:"address"`
}

func (h handler) UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	body := UpdateCustomerRequestBody{}
	body.FirstName = c.PostForm("first_name")
	body.LastName = c.PostForm("last_name")
	body.Email = c.PostForm("email")
	//body.Mobile = c.PostForm("mobile")
	body.Address = c.PostForm("address")
	//body.Password = c.PostForm("password")
	/*
		var FirstName = c.PostForm("first_name")
		var LastName = c.PostForm("last_name")
		var Email = c.PostForm("email")
		var Mobile = c.PostForm("mobile")
		var Address = c.PostForm("address")
		var Password = c.PostForm("password")
	*/
	// getting request's body
	/*
		if err := c.BindJSON(&body); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
	*/

	var customer models.Customer

	if result := h.DB.First(&customer, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	customer.FirstName = body.FirstName
	customer.LastName = body.LastName
	customer.Email = body.Email
	//customer.Mobile = body.Mobile
	customer.Address = body.Address
	result := h.DB.Save(&customer)
	if result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer data updated successfully", "data": &customer})
}
