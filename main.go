package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  uint32 `json:"price"`
}

// slice album go init
var albums = []album{
	{ID: "1", Title: "Dumes", Artist: "Denny Chaknan", Price: 100000},
	{ID: "2", Title: "Lamunan", Artist: "Shinta Arshinta", Price: 200000},
	{ID: "3", Title: "Kelangan", Artist: "Guyon Waton", Price: 250000},
}

func main() {
	router := gin.Default()

	// get all albums
	router.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)
	})

	// get albums by id
	router.GET("/albums/:id", getAlbumsById)

	// post new album
	router.POST("/albums", postAlbums)

	// run the http
	router.Run("localhost:8080")
}

func getAlbumsById(c *gin.Context) {
	id := c.Param("id")

	// search the id
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// bind the recieve json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// post newAlbum to albums
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
