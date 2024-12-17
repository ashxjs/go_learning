package application

import "go_learning/src/album/domain"

type AlbumService struct {
	repo domain.AlbumRepository
}

func NewAlbumService(repo domain.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) GetAlbums() ([]domain.Album, error) {
	return s.repo.GetAlbums()
}

func (s *AlbumService) GetAlbumByID(id string) (domain.Album, error) {
	return s.repo.GetAlbumByID(id)
}

func (s *AlbumService) CreateAlbum(album domain.Album) (domain.Album, error) {
	return s.repo.CreateAlbum(album)
}
