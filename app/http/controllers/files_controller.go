package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type FilesController struct {

}

func (FilesController) Store(c *gin.Context) {
	file, _ := c.FormFile("file")

	fileName := strconv.FormatInt(time.Now().UnixNano(), 10) + file.Filename

	c.SaveUploadedFile(file, "./storage/images/" + fileName)

	c.String(200, fileName)
}

