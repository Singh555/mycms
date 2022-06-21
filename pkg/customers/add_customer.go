package customers

import (
	"github.com/Singh555/mycms/common/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type AddCustomerRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}

func (h handler) AddCustomer(c *gin.Context) {
	body := AddCustomerRequestBody{}
	body.FirstName = c.PostForm("first_name")
	body.LastName = c.PostForm("last_name")
	body.Email = c.PostForm("email")
	body.Mobile = c.PostForm("mobile")
	body.Address = c.PostForm("address")
	body.Password = c.PostForm("password")
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
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "validation error", "error": err})
			return
		}
	*/
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 8)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while password hashing", "error": err})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while inserting data", "error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "customer data added successfully", "data": &customer})
}
