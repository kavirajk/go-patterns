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
}

type UserFriendsStore interface {
	Save(*models.UserFriends) error
}

type AlbumStore interface {
	Save(*models.Album) error
	Get(id uint) (*models.Album, error)
	GetBySlug(slug string) (*models.Album, error)
	GetAll() ([]models.Album, error)
}

type PictureStore interface {
	Save(*models.Picture) error
	Get(id uint) (*models.Picture, error)
	GetByAlbum(albumId uint) ([]models.Picture, error)
	GetAll() ([]models.Album, error)
}
