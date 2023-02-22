package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/yuhori/recipe-viewer/controllers"
)

func Init() {
	r := gin.Default()
	rc := controllers.NewRecipeController()
	v1 := r.Group("/api/v1/recipes")
	v1.GET("/list", rc.FetchAll)
	v1.GET("/:id", rc.Fetch)
	v1.GET("/search", rc.Search)
	v1.POST("/store", rc.Store)
	v1.POST("/delete", rc.Store)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
