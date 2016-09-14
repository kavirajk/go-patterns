package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type store struct {
	db      *gorm.DB
	user    UserStore
	album   AlbumStore
	picture PictureStore
}

func (s store) User() UserStore {
	return s.user
}

func (s store) Album() AlbumStore {
	return s.album
}

func (s store) Picture() PictureStore {
	return s.picture
}

func NewStore() *store {
	st := InitStore()
	st.user = NewUserStore(st)
	st.album = NewAlbumStore(st)
	st.picture = NewPictureStore(st)
	return st
}

func InitStore() *store {
	db, err := gorm.Open("postgres", "user=kaviraj password=kaviraj dbname=patterns sslmode=disable")
	if err != nil {
		panic(err)
	}
	st := &store{db: db}
	return st
}
