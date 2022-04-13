package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Product_uid string         `gorm:"index;type:varchar(22);primaryKey"`
	User_uid    string         `gorm:"index;type:varchar(22)"`
	Name        string
	Image       string `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
	Description string
	Price       string
	Stock       int
}
