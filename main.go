package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Tags          []string  `json:"tags"`
	Ingredients   []string  `json:"Ingredients"`
	Instructions  []string  `json:"Instructions"`
	DatePublished time.Time `json:"DatePublished"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, _ := ioutil.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

func updateRecipe(c *gin.Context) {
	var recipe Recipe
	id := c.Params.ByName("id")
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	foundAt := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			foundAt = i
			break
		}
	}
	if foundAt == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "recipe not found"})
		return
	}
	recipes[foundAt] = recipe
	c.JSON(http.StatusOK, recipe)
}

func createRecipe(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	guid := xid.New()
	recipe.ID = guid.String()
	recipe.DatePublished = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func getRecipes(c *gin.Context) {
	c.JSON(http.StatusOK, recipes)
}

func deleteRecipe(c *gin.Context) {
	id := c.Params.ByName("id")
	foundAt := -1
	for i := 0; i < len(recipes); i++ {
		if recipes[i].ID == id {
			foundAt = i
			break
		}
	}

	if foundAt == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "cannot delete recipe, not found"})
		return
	}

	recipes = append(recipes[:foundAt], recipes[foundAt+1:]...)
	c.JSON(http.StatusOK, gin.H{"message": "Recipe deleted successfully"})
}

func main() {
	router := gin.Default()
	router.POST("/recipe", createRecipe)
	router.GET("/recipes", getRecipes)
	router.PUT("/recipe/:id", updateRecipe)
	router.DELETE("/recipe/:id", deleteRecipe)
	router.Run()
}
