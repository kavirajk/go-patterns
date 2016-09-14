package store

import (
	"fmt"

	"github.com/kavirajk/go-patterns/models"
)

type pictureStore struct {
	*store
}

func NewPictureStore(st *store) *pictureStore {
	st.db.AutoMigrate(new(models.Picture))
	return &pictureStore{st}
}

func (ps *pictureStore) Save(pic *models.Picture) error {
	if err := pic.Valid(); err != nil {
		return err
	}
	if err := ps.db.Save(pic).Error; err != nil {
		return fmt.Errorf("picture.save: %s", err)
	}
	return nil
}

func (ps *pictureStore) get(pic *models.Picture, where ...interface{}) error {
	if err := ps.db.First(pic, where...).Error; err != nil {
		return err
	}
	return nil
}

func (ps *pictureStore) Get(id uint) (*models.Picture, error) {
	var pic models.Picture
	if err := ps.get(&pic, "id=?", id); err != nil {
		return nil, fmt.Errorf("picture.get.id: %d, %s", id, err)
	}
	return &pic, nil
}

func (ps *pictureStore) GetByAlbum(albumId uint) ([]models.Picture, error) {
	var pics []models.Picture
	if err := ps.db.Find(&pics, "album_id=?", albumId).Error; err != nil {
		return nil, fmt.Errorf("picture.get_by_album.id: %d, %s", albumId, err)
	}
	return pics, nil
}

func (ps *pictureStore) GetAll() ([]models.Picture, error) {
	var pics []models.Picture
	if err := ps.db.Find(&pics).Error; err != nil {
		return nil, fmt.Errorf("picture.get_all: %s", err)
	}
	return pics, nil
}