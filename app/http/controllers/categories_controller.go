package controllers

import (
	"buhang/app/models/category"
	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	BaseController
}

func (cc *CategoriesController) Index(c *gin.Context) {
	cc.success(c, category.List(), "success")
}

func (cc *CategoriesController) Create(c *gin.Context) {
	category.Create(c.Query("name"))
	cc.success(c, nil, "success")
}

func (cc *CategoriesController) Update(c *gin.Context) {
	category.Update(c.Query("id"), category.Category{Name: c.Query("name")})
	cc.success(c, nil, "success")
}
