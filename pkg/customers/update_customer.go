package customers

import (
	"github.com/Singh555/mycms/common/helper"
	"net/http"

	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
)

type UpdateCustomerRequestBody struct {
	Id        int64  `form:"id" binding:"required"`
	FirstName string `form:"first_name"`
	LastName  string `form:"last_name"`
	Email     string `form:"email" binding:"email"`
	//Mobile    string `json:"mobile"`
	Address string `form:"address"` //use json when request from postman is json data
}

func (h handler) UpdateCustomer(c *gin.Context) {

	body := UpdateCustomerRequestBody{}
	//body.FirstName = c.PostForm("first_name")
	//body.LastName = c.PostForm("last_name")
	//body.Email = c.PostForm("email")
	//body.Mobile = c.PostForm("mobile")
	//body.Address = c.PostForm("address")
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

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}
	id := body.Id
	var customer models.Customer

	if result := h.DB.First(&customer, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while getting customer data", "error": result.Error})
		return
	}

	customer.FirstName = body.FirstName
	customer.LastName = body.LastName
	customer.Email = body.Email
	//customer.Mobile = body.Mobile
	customer.Address = body.Address
	result := h.DB.Save(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while updating customer data", "error": result.Error})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer data updated successfully", "data": &customer})
}
