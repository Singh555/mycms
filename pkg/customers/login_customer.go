package customers

import (
	"fmt"
	"github.com/Singh555/mycms/common/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/Singh555/mycms/common/models"
	"net/http"
)

type TokenRequest struct {
	Mobile   string `form:"mobile" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h handler) LoginCustomer(context *gin.Context) {
	var request TokenRequest
	var user models.Customer
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	fmt.Println(request.Mobile)
	// check if email exists and password is correct
	record := h.DB.Order("id DESC").Where("mobile = ?", request.Mobile).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.GenerateJWT(user.Email, user.Mobile)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
