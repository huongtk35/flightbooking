package repositories

import (
	"flightbooking/db"
	"flightbooking/util"

	"gorm.io/gorm"
)

type IUserService interface {
	Create(user *db.User) (*db.User, error)
	Update(user *db.User) (*db.User, error)
	GetByUsername(username string) (*db.User, error)
	ValidatePassword(user *db.User, password string) bool
}

type UserService struct {
	IUserService
	DB *gorm.DB
}

func NewUserService() UserService {
	var db = db.NewGormDB()
	return UserService{DB: db}
}

func (u *UserService) Create(user *db.User) (*db.User, error) {
	create := u.DB.Create(user)
	if create.Error != nil {
		return user, create.Error
	}
	return user, nil
}

func (u *UserService) Update(user *db.User) (*db.User, error) {
	save := u.DB.Save(user)
	if save.Error != nil {
		return user, save.Error
	}
	return user, nil
}

func (u *UserService) GetByUsername(username string) (*db.User, error) {
	query := db.User{}
	find := u.DB.Find(&query, db.User{Username: username})
	if find.Error != nil {
		return nil, find.Error
	}
	return &query, nil
}

func (u *UserService) ValidatePassword(user *db.User, password string) bool {
	return util.ValidatePassword(password, user.Password)
}
