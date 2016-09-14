package store

import "github.com/kavirajk/go-patterns/models"

type Store interface {
	User() UserStore
	Album() AlbumStore
	Picture() PictureStore
}

type UserStore interface {
	Save(*models.User) error
	AddFriend(useId, friendId uint) error
	Get(id uint) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetAll() ([]models.User, error)
	Delete(id uint) error
	DeletePermanent(id uint) error
}

type AlbumStore interface {
	Save(*models.Album) error
	Get(id uint) (*models.Album, error)
	GetActive(id uint) (*models.Album, error)
	GetBySlug(slug string) (*models.Album, error)
	GetActiveBySlug(slug string) (*models.Album, error)
	GetByOwner(ownerId uint) ([]models.Album, error)
	GetActiveByOwner(ownerId uint) ([]models.Album, error)
	GetAll() ([]models.Album, error)
	GetAllActive() ([]models.Album, error)
	Delete(id uint) error
	DeletePermanent(id uint) error
}

type PictureStore interface {
	Save(*models.Picture) error
	Get(id uint) (*models.Picture, error)
	GetByAlbum(albumId uint) ([]models.Picture, error)
	GetAll() ([]models.Picture, error)
	Delete(id uint) error
	DeletePermanent(id uint) error
}
