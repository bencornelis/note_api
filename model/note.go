package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Text   string `json:"text"`
	Active bool   `json:"active"`

	UserID uint
}

func (s *DatabaseStore) GetNotes(userId uint) ([]Note, error) {
	var notes []Note
	user := User{}
	user.ID = userId
	if err := s.db.Model(&user).Related(&notes).Error; err != nil {
		return nil, err
	}
	fmt.Println(notes)
	return notes, nil
}

func (s *DatabaseStore) GetNote(id uint, userId uint) (Note, error) {
	var note Note
	if err := s.db.Find(&note, id).Error; err != nil {
		return Note{}, err
	}

	if note.UserID != userId {
		return Note{}, errors.New("note does not belong to user")
	}
	return note, nil
}

func (s *DatabaseStore) CreateNote(note Note) (Note, error) {
	note.Active = true
	if err := s.db.Create(&note).Error; err != nil {
		return Note{}, err
	}
	return note, nil
}

func (s *DatabaseStore) UpdateNote(id uint, note Note) (Note, error) {
	note.ID = id
	if err := s.db.Save(&note).Error; err != nil {
		return Note{}, err
	}
	return note, nil
}

func (s *DatabaseStore) DeleteNote(id uint, userId uint) error {
	note, err := s.GetNote(id, userId)
	if err != nil {
		return err
	}

	return s.db.Delete(&note).Error
}
