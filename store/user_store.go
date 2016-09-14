package store

import (
	"fmt"

	"github.com/kavirajk/go-patterns/models"
)

type userStore struct {
	*store
}

func NewUserStore(st *store) *userStore {
	st.db.AutoMigrate(new(models.User), new(models.UserFriends))
	return &userStore{st}
}

func (us *userStore) Save(user *models.User) error {
	if err := user.Valid(); err != nil {
		return err
	}
	if err := us.db.Save(user).Error; err != nil {
		return fmt.Errorf("user.save: %s", err)
	}
	return nil
}

func (us *userStore) AddFriend(userId, friendId uint) error {
	friend := models.UserFriends{userId, friendId}
	if err := us.db.Save(&friend).Error; err != nil {
		return fmt.Errorf("user.add_friend: %s", err)
	}
	return nil
}

func (us *userStore) Get(id uint) (*models.User, error) {
	var u models.User
	if err := us.get(&u, "id=?", id); err != nil {
		return nil, fmt.Errorf("user.get: %d, %s", id, err)
	}

	return &u, nil
}

func (us *userStore) GetByUsername(username string) (*models.User, error) {
	var u models.User
	if err := us.get(&u, "username=?", username); err != nil {
		return nil, fmt.Errorf("user.get.byusername: %s, %s", username, err)
	}
	return &u, nil
}

func (us *userStore) GetAll() ([]models.User, error) {
	var users []models.User
	if err := us.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("user.get.all: %s", err)
	}
	return users, nil
}

func (us *userStore) get(u *models.User, where ...interface{}) error {
	if err := us.db.First(u, where...).Error; err != nil {
		return err
	}
	if err := us.db.Table("users").Joins("inner join user_friends on user_friends.friend_id=users.id").Where("user_friends.user_id=?", u.Id).Find(&u.Friends).Error; err != nil {
		return err
	}
	return nil
}

func (us *userStore) Delete(id uint) error {
	user, err := us.Get(id)
	if err != nil {
		return err
	}

	albums, err := us.album.GetByOwner(id)
	if err != nil {
		return err
	}

	for i := range albums {
		if err := us.album.Delete(albums[i].Id); err != nil {
			return err
		}
	}

	if err := us.db.Delete(user).Error; err != nil {
		return fmt.Errorf("user.delete.id: %d, %s", id, err)
	}
	return nil
}

func (us *userStore) DeletePermanent(id uint) error {
	user, err := us.Get(id)
	if err != nil {
		return err
	}

	albums, err := us.album.GetByOwner(id)
	if err != nil {
		return err
	}

	for i := range albums {
		if err := us.album.DeletePermanent(albums[i].Id); err != nil {
			return err
		}
	}

	if err := us.db.Unscoped().Delete(user).Error; err != nil {
		return fmt.Errorf("user.delete_permanent.id: %d, %s", id, err)
	}
	return nil
}
