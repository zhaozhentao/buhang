package routes

import (
	"buhang/app/http/controllers"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func Init() {
	R = gin.Default()

	cc := controllers.CardsController{}
	R.POST("/api/cards", cc.Create)
	R.GET("/api/cards", cc.Index)
	R.GET("/api/cards/:id", cc.Show)
	R.DELETE("/api/cards/:id", cc.Delete)

	fc := controllers.FilesController{}
	R.POST("/api/files", fc.Store)

	cgc := controllers.CategoriesController{}
	R.GET("/api/categories", cgc.Index)
	R.POST("/api/categories", cgc.Create)
	R.PATCH("/api/categories/:id", cgc.Update)
	R.DELETE("/api/categories/:id", cgc.Delete)
}
