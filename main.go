package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
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

func searchRecipes(c *gin.Context) {
	tag := c.Query("tag")
	searchResult := make([]Recipe, 0)
	found := false
	for i := 0; i < len(recipes); i++ {
		for _, t := range recipes[i].Tags {
			if strings.EqualFold(tag, t) {
				searchResult = append(searchResult, recipes[i])
				found = true
				break
			}

		}
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "no recipe was found"})
	} else {
		c.JSON(http.StatusOK, searchResult)
	}
}

func main() {
	router := gin.Default()
	router.POST("/recipe", createRecipe)
	router.GET("/recipes", getRecipes)
	router.GET("/recipes/search", searchRecipes)
	router.PUT("/recipe/:id", updateRecipe)
	router.DELETE("/recipe/:id", deleteRecipe)
	router.Run()
}
