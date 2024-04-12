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
	router.GET("/albums", getAlbums)

	// run the http
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
