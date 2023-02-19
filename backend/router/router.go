package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/yuhori/recipe-viewer/controllers"
)

func Init() {
	r := gin.Default()
	r.GET("/api/v1/list", controllers.FetchAll)
	r.GET("/api/v1/:id", controllers.Fetch)
	r.GET("/api/v1/search", controllers.Search)
	r.POST("/api/v1", controllers.Store)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
