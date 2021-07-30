package controllers

import (
	"buhang/app/models/card"
	"buhang/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardsController struct {
	BaseController
}

func (controller *CardsController) Create(c *gin.Context) {
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

	controller.success(c, nil, "success")
}

func (controller *CardsController) Index(c *gin.Context) {
	controller.success(c, card.Index(), "success")
}

func (*CardsController) Show(c *gin.Context) {
	c.String(http.StatusOK, "Show")
}

func (*CardsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}
