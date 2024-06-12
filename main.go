package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"rsc.io/quote"
)

type album struct {
	id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	port := os.Getenv("GO_PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/albums", getAlbums)

	router.Run(":" + port)
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
