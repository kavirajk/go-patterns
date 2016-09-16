package store

import (
	"fmt"
	"log"

	"github.com/kavirajk/go-patterns/models"
)

type pictureStore struct {
	*store
}

func NewPictureStore(st *store) *pictureStore {
	ps := &pictureStore{st}
	ps.CreateTableIfNotExist()
	ps.CreateIndexesIfNotExists()
	return ps
}

func (ps *pictureStore) CreateTableIfNotExist() {
	if !ps.db.HasTable(&models.Picture{}) {
		ps.db.CreateTable(&models.Picture{})
		ps.db.Model(&models.Picture{}).AddForeignKey("album_id", "albums(id)", "RESTRICT", "RESTRICT")
		if ps.db.Error != nil {
			log.Fatalf("critical.picture.migrate.create_table: %s", ps.db.Error)
		}
	}
}

func (ps *pictureStore) CreateIndexesIfNotExists() {
	indexes := map[string]string{
		"idx_picture_created_at": "created_at",
		"idx_picture_updated_at": "updated_at",
		"idx_picture_deleted_at": "deleted_at",
	}
	for k, v := range indexes {
		ps.db.Model(&models.Picture{}).AddIndex(k, v)
		if ps.db.Error != nil {
			log.Fatalf("critical.picture.migrate.create_indexes: %s", ps.db.Error)
		}
	}
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

func (ps *pictureStore) Delete(id uint) error {
	pic, err := ps.Get(id)
	if err != nil {
		return err
	}
	if err := ps.db.Delete(pic).Error; err != nil {
		return fmt.Errorf("picture.delete.id: %d, %s", id, err)
	}
	return nil
}

func (ps *pictureStore) DeletePermanent(id uint) error {
	pic, err := ps.Get(id)
	if err != nil {
		return err
	}
	if err := ps.db.Unscoped().Delete(pic).Error; err != nil {
		return fmt.Errorf("picture.delete_permanent.id: %d, %s", id, err)
	}
	return nil
}
