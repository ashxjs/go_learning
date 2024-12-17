package domain

type AlbumRepository interface {
	GetAlbums() ([]Album, error)
	GetAlbumByID(id string) (Album, error)
	CreateAlbum(album Album) (Album, error)
}
