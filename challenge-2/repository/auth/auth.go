package auth

import (
	"be/entities"
	"be/utils"
	"errors"

	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Login(userName, password string) (string, error) {

	var user entities.User

	if res := r.db.Model(&entities.User{}).Where("user_name = ?", userName).Find(&user); res.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("incorrect password")
	}

	return user.User_uid, nil
}
