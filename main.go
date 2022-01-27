package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	FirstName string `xml:"FirstName,attr"`
	LastName  string `xml:"LastName,attr"`
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello World",
	})
}
func person(c *gin.Context) {
	c.XML(200, Person{FirstName: "Youssef", LastName: "Boubekri"})
}
func content(c *gin.Context) {
	param := c.Params.ByName("type")
	if param == "xml" {
		c.XML(200, Person{FirstName: "Youssef", LastName: "Boubekri"})

	} else {
		c.JSON(200, Person{FirstName: "Youssef", LastName: "Boubekri"})

	}
}
func params(c *gin.Context) {
	param := c.Params.ByName("param")
	c.JSON(200, gin.H{
		"message": "hello " + param,
	})
}
func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/person", person)
	router.GET("/person/:type", content)
	router.GET("/:param", params)
	router.Run("127.0.0.1:8080")
}
