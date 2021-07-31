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
		CategoryId:   types.StringToInt(c.Query("category_id")),
		GoodsNumber:  c.Query("goods_number"),
		GoodsName:    c.Query("goods_name"),
		Weight:       c.Query("weight"),
		Width:        types.StringToFloat(c.Query("width")),
		Img:          c.Query("img"),
		MeterPerKilo: types.StringToFloat(c.Query("meterPerKilo")),
	}

	_card.Create()

	controller.success(c, nil, "success")
}

func (controller *CardsController) Index(c *gin.Context) {
	controller.success(c, card.Index(), "success")
}

func (controller *CardsController) Show(c *gin.Context) {
	controller.success(c, card.Show(c.Param("id")), "success")
}

func (*CardsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}
