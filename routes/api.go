package routes

import (
	"buhang/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	cc := controllers.CardsController{}
	r.POST("/api/cards", cc.Create)
	r.GET("/api/cards", cc.Index)
	r.GET("/api/cards/:id", cc.Show)
	r.DELETE("/api/cards/:id", cc.Delete)

	fc := controllers.FilesController{}
	r.POST("/api/files", fc.Store)

	cgc := controllers.CategoriesController{}
	r.GET("/api/categories", cgc.Index)
	r.POST("/api/categories", cgc.Create)
	r.PATCH("/api/categories/:id", cgc.Update)

	r.Static("/manager", "./public/manager")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
