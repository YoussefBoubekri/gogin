package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Recipe struct {
	Name          string    `json:"Name"`
	Tags          []string  `json:"Tags"`
	Ingredients   []string  `json:"Ingredients"`
	Instructions  []string  `json:"Instructions"`
	DatePublished time.Time `json:"DatePublished"`
}

func main() {
	router := gin.Default()
	router.Run()
}
