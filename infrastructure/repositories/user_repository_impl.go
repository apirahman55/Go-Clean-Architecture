package repositories

import (
	"errors"
	"user-services/domain/entities"
	"user-services/domain/repository"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{db}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r *UserRepositoryImpl) FindAll() (users *entities.Users, err error) {
	err = r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) FindOne(uuid string) (user *entities.User, err error) {
	err = r.db.First(&user, uuid).Error
	if err != nil {
		return nil, err
	}

	if user.ID.String() == "" {
		return nil, errors.New("cannot finding user with" + uuid)
	}

	return user, nil
}

func (r *UserRepositoryImpl) Insert(user *entities.User) (err error) {
	err = r.db.Save(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) Delete(uuid string) (err error) {
	return err
}
