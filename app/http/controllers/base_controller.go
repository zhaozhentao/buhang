package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
}

func (BaseController) success(c *gin.Context, data interface{}, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": data,
		"message": message,
	})
}
