// adapter/http_album_handler.go
package adapter

import (
	"learning_session/src/album/application"
	"learning_session/src/album/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	service *application.AlbumService
}

func NewAlbumHandler(service *application.AlbumService) *AlbumHandler {
	return &AlbumHandler{service: service}
}

func (h *AlbumHandler) GetAlbums(c *gin.Context) {
	albums, err := h.service.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := h.service.GetAlbumByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
		return
	}
	c.JSON(http.StatusOK, album)
}

func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
	var newAlbum domain.Album
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	createdAlbum, err := h.service.CreateAlbum(newAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdAlbum)
}
