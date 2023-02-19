package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/yuhori/recipe-viewer/controllers"
)

func Init() {
	r := gin.Default()
	r.GET("/list", controllers.FetchAll)
	r.GET("/:id", controllers.Fetch)
	r.GET("/search", controllers.Search)
	r.POST("/", controllers.Store)
	r.Group("/api/v1")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
