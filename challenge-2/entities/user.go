package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	User_uid  string         `gorm:"index;type:varchar(22);primaryKey"`
	UserName  string         `gorm:"index;not null;type:varchar(100)"`
	Email     string         `gorm:"index;not null;type:varchar(100)"`
	Password  string         `gorm:"not null;type:varchar(100)"`
	Name      string
	Image     string    `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
	Products  []Product `gorm:"foreignKey:user_uid;references:user_uid"`
}
