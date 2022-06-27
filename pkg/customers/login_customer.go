package customers

import (
	"fmt"
	"github.com/Singh555/mycms/common/auth"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"github.com/Singh555/mycms/common/models"
	"net/http"
)

//structure to parse login request body
type TokenRequest struct {
	Mobile   string `form:"mobile" binding:"required"`
	Password string `form:"password" binding:"required"`
}

//function to match login credentials and generate jwt tokens
func (h handler) LoginCustomer(context *gin.Context) {
	var request TokenRequest
	var user models.Customer
	if err := context.ShouldBind(&request); err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	fmt.Println(request.Mobile)
	// check if customer exists
	record := h.DB.Order("id DESC").Where("mobile = ?", request.Mobile).First(&user)
	if record.Error != nil {
		log.Error(record.Error)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": record.Error})
		context.Abort()
		return
	}
	// check if password is correct
	credentialError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if credentialError != nil {
		log.Error(credentialError.Error())
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.ID, user.Mobile, user.Email)
	if err != nil {
		log.Error(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
