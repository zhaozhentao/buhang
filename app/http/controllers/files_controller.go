package controllers

import (
	"buhang/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type FilesController struct {
}

func (FilesController) Store(c *gin.Context) {
	file, _ := c.FormFile("file")

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano()/1e6, file.Filename)

	c.SaveUploadedFile(file, config.Viper.GetString("UPLOAD")+fileName)

	c.String(200, fileName)
}
