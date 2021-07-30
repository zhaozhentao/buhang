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
	_card := card.Card{
		CategoryId:   types.StringToInt(c.PostForm("category_id")),
		GoodsNumber:  c.PostForm("goods_number"),
		GoodsName:    c.PostForm("goods_name"),
		Weight:       c.PostForm("weight"),
		Width:        types.StringToFloat(c.PostForm("width")),
		Img:          c.PostForm("img"),
		MeterPerKilo: types.StringToFloat(c.PostForm("meterPerKilo")),
	}

	_card.Create()

	c.JSON(200, gin.H{
		"code":    20000,
		"data":    nil,
		"message": "success",
	})
}

func (*CardsController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": card.Index(),
	})
}

func (*CardsController) Show(c *gin.Context) {
	c.String(http.StatusOK, "Show")
}

func (*CardsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}
