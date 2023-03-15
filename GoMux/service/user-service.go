package service

import (
	"khushal/goMux/entity"

	"gorm.io/gorm"
)

type UserService interface {
	GetUsers() []entity.User
	GetUser(id uint) entity.User
	CreateUser(user entity.User)
	UpdateUser(user entity.User)
	DeleteUser(id uint)
	GetUserByEmail(email string) entity.User
}

type service struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserService{
	return &service{db}
}

func (service service) GetUsers() []entity.User{
	var users []entity.User
	service.db.Find(&users)
	return users
}

func (service service) GetUser(id uint) entity.User{
	var user entity.User
	service.db.First(&user, id)
	return user
}

func (service service) UpdateUser(user entity.User) {
	service.db.Save(&user)
}

func (service service) CreateUser(user entity.User) {
	service.db.Create(&user)
}

func (service service) DeleteUser(id uint) {
	var user entity.User
	service.db.Delete(&user, id)
}

func (service service) GetUserByEmail(email string) entity.User{
	var user entity.User
	service.db.Where("email = ?", email).First(&user)
	return user
}
