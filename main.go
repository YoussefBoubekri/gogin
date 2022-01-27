package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID            string    `json:"id"`
	Name          string    `json:"Name"`
	Tags          []string  `json:"Tags"`
	Ingredients   []string  `json:"Ingredients"`
	Instructions  []string  `json:"Instructions"`
	DatePublished time.Time `json:"DatePublished"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

func createRecipe(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	guid := xid.New()
	recipe.ID = guid.String()
	recipe.DatePublished = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()
	router.POST("/recipe", createRecipe)
	router.Run()
}
