package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`

	Notes []Note
}

func (s *DatabaseStore) CreateUser(user User) (User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return User{}, err
	}
	user.Password = string(hashedPassword)

	if err := s.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (s *DatabaseStore) FindUser(username string) (User, error) {
	var user User
	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
