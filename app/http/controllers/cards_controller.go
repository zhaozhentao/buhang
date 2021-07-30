package controllers

import (
	"buhang/app/models/card"
	"buhang/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardsController struct {
}

func (*CardsController) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"categoryId":   types.StringToInt(c.PostForm("category_id")),
		"goodsNumber":  c.PostForm("goods_number"),
		"goods_name":   c.PostForm("goods_name"),
		"weight":       c.PostForm("weight"),
		"width":        c.PostForm("width"),
		"img":          c.PostForm("img"),
		"meterPerKilo": c.PostForm("meterPerKilo"),
	})
}

func (*CardsController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, card.Index())
}

func (*CardsController) Show(c *gin.Context) {
	c.String(http.StatusOK, "Show")
}

func (*CardsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}
