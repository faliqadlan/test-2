package entities

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Movie_uid  string `gorm:"index;type:varchar(22);primaryKey"`
	Title       string `gorm:"type:varchar(255)"`
	Description string
	Duration    string
	Artist      string
	Genres      string
	Image       string `gorm:"default:'https://www.teralogistics.com/wp-content/uploads/2020/12/default.png'"`
}
