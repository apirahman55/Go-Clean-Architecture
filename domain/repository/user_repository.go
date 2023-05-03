package repository

import "user-services/domain/entities"

type UserRepository interface {
	FindAll() (users *entities.Users, err error)
	FindOne(uuid string) (user *entities.User, err error)
	Insert(user *entities.User) (err error)
	Delete(uuid string) (err error)
}
