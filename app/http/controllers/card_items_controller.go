package controllers

import (
	"buhang/app/models/card_item"
	"github.com/gin-gonic/gin"
	"strings"
)

type CardItemsController struct {
	BaseController
}

func (cc *CardItemsController) Create(c *gin.Context) {
	images := strings.Split(c.Query("image"), ",")

	card_item.Create(images, c.Query("cardId"))

	cc.success(c, nil, "success")
}

func (cc *CardItemsController) Delete(c *gin.Context) {
	card_item.Delete(c.Param("id"))
	cc.success(c, nil, "success")
}
