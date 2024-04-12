package main

import (
	"fmt"
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
		c.IndentedJSON(http.StatusOK, gin.H{"data": albums, "message": "success retrieve albums"})
	})

	// get albums by id
	router.GET("/albums/:id", getAlbumById)

	// post new album
	router.POST("/albums", postAlbums)

	// update album by id
	router.PUT("/albums/:id", updateAlbumById)

	// delete album by id
	router.DELETE("/albums/:id", deleteAlbumById)

	// run the http
	router.Run("localhost:8080")
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	// search the id
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": album, "message": "success retrieve album"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	newAlbum.ID = fmt.Sprint(len(albums))

	// bind the recieve json to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// post newAlbum to albums
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, gin.H{"data": newAlbum, "message": "success created album"})
}

func updateAlbumById(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	// bind the recieve json to album
	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	if updatedAlbum.ID == "" || updatedAlbum.ID != id {
		return
	}

	for i := 0; i < len(albums); i++ {
		if albums[i].ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, gin.H{"data: ": updatedAlbum, "message": "success updated album"})
		}
	}
}

func deleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(albums); i++ {
		if albums[i].ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": albums[i], "message": "album has been deleted"})
			albums = append(albums[:i], albums[i+1:]...)
			return
		}
	}
}
