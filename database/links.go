package database

import (
	"gorm.io/gorm"
)

type Links struct {
	gorm.Model
	Title   string  `gorm:"size:255;column:Title"`
	Address string  `gorm:"size:255;column:Address"`
	UserID  *uint64 `gorm:"index;column:UserID"`
	Users   *Users  `gorm:"foreignKey:UserID"`
}

func (l *Links) GetAll(DB *gorm.DB, result *[]Links, take int, skip int) error {
	res := DB.Limit(take).Find(&result)

	if res.Error != nil {
		return res.Error
	}

	return nil
}
