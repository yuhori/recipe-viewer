package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/yuhori/recipe-viewer/controllers"
)

func Init() {
	r := gin.Default()

	// cors 対応
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://yuhori.github.io/recipe-viewer",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"DETETE",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	rc := controllers.NewRecipeController()
	v1 := r.Group("/api/v1/recipes")
	v1.GET("/list", rc.FetchAll)
	v1.GET("/:id", rc.Fetch)
	v1.GET("/search", rc.Search)
	v1.POST("/store", rc.Store)
	v1.POST("/delete", rc.Store)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
