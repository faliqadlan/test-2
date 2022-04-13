package user

import (
	"be/entities"
	"be/utils"
	"errors"
	"strconv"

	"github.com/google/uuid"
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

func (r *Repo) Create(req entities.User) (entities.User, error) {
	var uid string

	for {
		uid = strconv.Itoa(int(uuid.New().ID()))
		var res = r.db.Model(&entities.User{}).Where("user_uid = ?", uid).Scan(&entities.User{})
		if res.RowsAffected == 0 {
			break
		}
	}

	// check username

	var res = r.db.Model(&entities.User{}).Where("user_name = ?", req.UserName).Scan(&entities.User{})

	if res.RowsAffected != 0 {
		return entities.User{}, errors.New("user name is already exist")
	}

	// check email

	res = r.db.Model(&entities.User{}).Where("email = ?", req.Email).Scan(&entities.User{})

	if res.RowsAffected != 0 {
		return entities.User{}, errors.New("email is already exist")
	}

	// hash password
	var err error
	req.Password, err = utils.HashPassword(req.Password)
	
	if err != nil {
		return entities.User{}, err
	}

	req.User_uid = uid

	if res := r.db.Model(&entities.User{}).Create(&req); res.Error != nil {
		return entities.User{}, res.Error
	}

	return req, nil
}
