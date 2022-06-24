//A package to provide common functionality over the application
package helper

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"

	"mime/multipart"
	"path/filepath"
)

// function to parse validation error and send to response
func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//function to upload image using gin context

func UploadImageWithGin(c *gin.Context, file *multipart.FileHeader, directory string) string {

	// Create the folder if it doesn't
	// already exist
	err1 := os.MkdirAll(".."+directory, os.ModePerm)
	if err1 != nil {
		log.Error(err1.Error())

		return "error"
	}

	var extention = filepath.Ext(file.Filename)
	name := uuid.NewString() //using google uuid package to generate unique image name
	var storagePath = directory + name + extention

	log.Debug("Starting file to upload " + storagePath)

	err := c.SaveUploadedFile(file, ".."+storagePath)

	if err != nil {
		log.Error(err.Error())

		return "error"
	}

	return storagePath
}

// function to check if file exists
func DoesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		fmt.Printf("%v file does not exist\n", fileName)
		return false
	} else {
		fmt.Printf("%v file exist\n", fileName)
		return true
	}
	return false
}
