package database

import (
	"demographql/graph/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Username string `gorm:"size:127;column:Username"`
	Password string `gorm:"size:127;column:Password"`
}

func (user *Users) Create(DB *gorm.DB, input model.NewUser) {
	existUser := Users{}
	// find user
	DB.Model(&Users{}).First(&existUser)

}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
