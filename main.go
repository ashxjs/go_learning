// main.go
package main

import (
	"go_learning/src/album/adapter"
	"go_learning/src/album/application"
	"go_learning/src/album/infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := infrastructure.NewMemoryAlbumRepository()
	service := application.NewAlbumService(repo)
	handler := adapter.NewAlbumHandler(service)

	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumByID)
	router.POST("/albums", handler.CreateAlbum)
	router.Run(":8080")
}
