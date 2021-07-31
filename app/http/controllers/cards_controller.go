package controllers

import (
	"buhang/app/models/card"
	"buhang/app/models/card_item"
	"buhang/pkg/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CardsController struct {
	BaseController
}

func (cc *CardsController) Create(c *gin.Context) {
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

	cc.success(c, nil, "success")
}

func (cc *CardsController) Index(c *gin.Context) {
	cc.success(c, card.Index(), "success")
}

func (cc *CardsController) Show(c *gin.Context) {
	cc.success(c, card.Show(c.Param("id")), "success")
}

func (*CardsController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "delete")
}

func (cc *CardsController) Items(c *gin.Context) {
	cc.success(c, card_item.Items(c.Param("id")), "success")
}
