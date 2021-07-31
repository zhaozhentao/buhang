package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type FilesController struct {
}

func (FilesController) Store(c *gin.Context) {
	file, _ := c.FormFile("file")

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), file.Filename)

	c.SaveUploadedFile(file, "./storage/images/"+fileName)

	c.String(200, fileName)
}
