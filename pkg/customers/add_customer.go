package customers

import (
	"fmt"
	"github.com/Singh555/mycms/common/helper"
	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AddCustomerRequestBody struct {
	FirstName string `form:"first_name" binding:"required"`
	LastName  string `form:"last_name" binding:"required"`
	Email     string `form:"email" binding:"email"`
	Mobile    string `form:"mobile" binding:"required,min=10,max=12"`
	Address   string `form:"address" binding:"required"` //use json when request from postman is json data
	Password  string `form:"password" binding:"required,min=8"`
}

func (h handler) AddCustomer(c *gin.Context) {
	body := AddCustomerRequestBody{}
	/*
		body.FirstName = c.PostForm("first_name")
		body.LastName = c.PostForm("last_name")
		body.Email = c.PostForm("email")
		body.Mobile = c.PostForm("mobile")
		body.Address = c.PostForm("address")
		body.Password = c.PostForm("password")
	*/
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
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, helper.ErrorResponse(err))
		return
	}

	var checkCustomer models.Customer
	var count int64
	if err := h.DB.Model(&checkCustomer).Select("id").Where("mobile", body.Mobile).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err))
		return
	}
	if count > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Customer Already exists", "error": ""})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.ErrorResponse(err))
		return
	}
	var customer models.Customer

	customer.FirstName = body.FirstName
	customer.LastName = body.LastName
	customer.Email = body.Email
	customer.Mobile = body.Mobile
	customer.Address = body.Address
	customer.Password = string(hashedPassword)

	if result := h.DB.Create(&customer); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while inserting customer data", "error": result.Error})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer data added successfully", "data": &customer})
}
