package router

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/yuhori/recipe-viewer/controllers"
)

func Init() {
	r := gin.Default()
	r.GET("/", controllers.FetchAll)
	r.GET("/", controllers.Fetch)
	r.GET("/", controllers.Search)
	r.GET("/", controllers.Store)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
