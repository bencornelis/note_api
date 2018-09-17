package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseStore struct {
	db *gorm.DB
}

func NewDatabaseStore() (*DatabaseStore, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=bencornelis dbname=bencornelis sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Note{})
	db.AutoMigrate(&User{})
	s := &DatabaseStore{db}
	return s, nil
}

func (s *DatabaseStore) Close() {
	s.db.Close()
}
