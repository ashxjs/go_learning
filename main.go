// main.go
package main

import (
	"learning_session/src/album/adapter"
	"learning_session/src/album/application"
	"learning_session/src/album/infrastructure"

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
