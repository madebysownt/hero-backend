package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func UploadSingleFile(c *gin.Context) {
	// single file
	file, err := c.FormFile("file")
	if err != nil {
		return
	}

	// Upload the file to specific dst.
	dst := os.Getenv("HOME") + "/.hero/files/" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func UploadMultipleFile(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["file[]"]

	for _, file := range files {
		log.Println(file.Filename)

		// Upload the file to specific dst.
		dst := os.Getenv("HOME") + "/.hero/files/" + file.Filename
		if err := c.SaveUploadedFile(file, dst); err != nil {
			return
		}
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}
