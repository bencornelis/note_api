package model

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Text   string `json:"text"`
	Active bool   `json:"active"`
}

func (s *DatabaseStore) GetNotes() ([]Note, error) {
	var notes []Note
	if err := s.db.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func (s *DatabaseStore) GetNote(id uint) (Note, error) {
	var note Note
	if err := s.db.Find(&note, id).Error; err != nil {
		return Note{}, err
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

func (s *DatabaseStore) DeleteNote(id uint) error {
	note := Note{}
	note.ID = id
	return s.db.Delete(&note).Error
}
