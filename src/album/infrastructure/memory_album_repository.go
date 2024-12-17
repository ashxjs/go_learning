package infrastructure

import (
	"errors"
	"learning_session/src/album/domain"
)

type MemoryAlbumRepository struct {
	albums []domain.Album
}

func NewMemoryAlbumRepository() *MemoryAlbumRepository {
	return &MemoryAlbumRepository{
		albums: []domain.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func (r *MemoryAlbumRepository) GetAlbums() ([]domain.Album, error) {
	return r.albums, nil
}

func (r *MemoryAlbumRepository) GetAlbumByID(id string) (domain.Album, error) {
	for _, a := range r.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return domain.Album{}, errors.New("album not found")
}

func (r *MemoryAlbumRepository) CreateAlbum(album domain.Album) (domain.Album, error) {
	r.albums = append(r.albums, album)
	return album, nil
}
