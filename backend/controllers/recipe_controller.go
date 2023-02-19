package controllers

import (
	"encoding/json"
	"io"
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
		return
	}
	c.JSON(http.StatusOK, rs)
}

func (rc *RecipeController) Fetch(c *gin.Context) {
	r, err := models.GetOne(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"FetachOne": "NG",
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func (rc *RecipeController) Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Search": "OK",
	})
}

func (rc *RecipeController) Store(c *gin.Context) {
	var r models.Recipe
	if err := Parse(&r, c.Request.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Parse": "NG",
		})
		return
	}
	if err := models.Create(r); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Create": "NG",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Store": "OK",
	})
}

func (rc *RecipeController) Delete(c *gin.Context) {
	if err := models.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Delete": "NG",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Search": "OK",
	})
}

func Parse(out interface{}, stream io.ReadCloser) (err error) {
	var jsonData []byte
	if jsonData, err = io.ReadAll(stream); err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, out); err != nil {
		return err
	}
	return nil
}
