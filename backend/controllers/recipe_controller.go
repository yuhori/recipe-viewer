package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhori/recipe-viewer/models"
)

type RecipeController struct{}

func NewRecipeController() *RecipeController {
	if err := models.InitDB(); err != nil {
		panic(err)
	}
	return &RecipeController{}
}

func (rc *RecipeController) FetchAll(c *gin.Context) {
	rs, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"FetachAll": "NG",
		})
	}
	c.JSON(http.StatusOK, rs)
}

func (rc *RecipeController) Fetch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (rc *RecipeController) Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (rc *RecipeController) Store(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
