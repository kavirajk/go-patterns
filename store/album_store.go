package store

import (
	"fmt"
	"strings"

	"github.com/kavirajk/go-patterns/models"
)

type albumStore struct {
	*store
}

func NewAlbumStore(st *store) *albumStore {
	st.db.AutoMigrate(new(models.Album))
	return &albumStore{st}
}

func (as *albumStore) Save(album *models.Album) error {
	if err := album.Valid(); err != nil {
		return err
	}
	if err := as.db.Save(album).Error; err != nil {
		return fmt.Errorf("album.save: %s", err)
	}
	return nil
}

func (as *albumStore) getAlbum(album *models.Album, where ...interface{}) error {
	var err error
	if err = as.db.First(album, where...).Error; err != nil {
		return err
	}
	album.Pictures, err = as.picture.GetByAlbum(album.Id)
	if err != nil {
		return err
	}
	album.Owner, err = as.user.Get(album.OwnerId)
	if err != nil {
		return err
	}
	return nil
}

func (as *albumStore) Get(id uint) (*models.Album, error) {
	var a models.Album
	if err := as.getAlbum(&a, "id=?", id); err != nil {
		return nil, fmt.Errorf("album.get.id: %d %s", id, err)
	}
	return &a, nil
}

func (as *albumStore) GetActive(id uint) (*models.Album, error) {
	var a models.Album
	if err := as.getAlbum(&a, "id=? and is_active=?", id, true); err != nil {
		return nil, fmt.Errorf("album.get.id: %d %s", id, err)
	}
	return &a, nil
}

func (as *albumStore) GetBySlug(slug string) (*models.Album, error) {
	slug = strings.ToLower(slug)
	var a models.Album
	if err := as.getAlbum(&a, "slug=?", slug); err != nil {
		return nil, fmt.Errorf("album.get.slug: %s %s", slug, err)
	}
	return &a, nil
}

func (as *albumStore) GetActiveBySlug(slug string) (*models.Album, error) {
	slug = strings.ToLower(slug)
	var a models.Album
	if err := as.getAlbum(&a, "slug=? and is_active=?", slug, true); err != nil {
		return nil, fmt.Errorf("album.get.slug: %s %s", slug, err)
	}
	return &a, nil
}

func (as *albumStore) GetByOwner(ownerId uint) ([]models.Album, error) {
	var albums []models.Album
	if err := as.db.Find(&albums, "owner_id=?", ownerId).Error; err != nil {
		return nil, fmt.Errorf("album.get_by_owner.id: %d, %s", ownerId, err)
	}
	return albums, nil
}

func (as *albumStore) GetActiveByOwner(ownerId uint) ([]models.Album, error) {
	var albums []models.Album
	if err := as.db.Find(&albums, "owner_id=? and is_active=?", ownerId, true).Error; err != nil {
		return nil, fmt.Errorf("album.get_by_owner.id: %d, %s", ownerId, err)
	}
	return albums, nil
}

func (as *albumStore) GetAll() ([]models.Album, error) {
	var albums []models.Album
	if err := as.db.Find(&albums).Error; err != nil {
		return nil, fmt.Errorf("album.get_all: %s", err)
	}
	return albums, nil
}

func (as *albumStore) GetAllActive() ([]models.Album, error) {
	var albums []models.Album
	if err := as.db.Find(&albums, "is_active=?", true).Error; err != nil {
		return nil, fmt.Errorf("album.get_all: %s", err)
	}
	return albums, nil
}

func (as *albumStore) Delete(id uint) error {
	album, err := as.Get(id)
	if err != nil {
		return err
	}
	for i := range album.Pictures {
		if err := as.picture.Delete(album.Pictures[i].Id); err != nil {
			return err
		}
	}
	if err := as.db.Delete(album).Error; err != nil {
		return fmt.Errorf("picture.delete.id: %d, %s", id, err)
	}
	return nil
}

func (as *albumStore) DeletePermanent(id uint) error {
	album, err := as.Get(id)
	if err != nil {
		return err
	}
	for i := range album.Pictures {
		if err := as.picture.DeletePermanent(album.Pictures[i].Id); err != nil {
			return err
		}
	}
	if err := as.db.Unscoped().Delete(album).Error; err != nil {
		return fmt.Errorf("picture.delete_permanent.id: %d, %s", id, err)
	}
	return nil
}
