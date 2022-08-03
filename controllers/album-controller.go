package controllers

import (
	"net/http"

	"github.com/akposieyefa/models"
	"github.com/gin-gonic/gin"
)

//get all albums
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

//crete album
func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

//get album by id
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	albums := models.Albums

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
