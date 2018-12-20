package data

import (
	"github.com/jinzhu/gorm"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Participant struct
type Participant struct {
	gorm.Model
	Surname      string
	Name         string
	Nomination   Nomination `gorm:"foreignkey:NominationId"`
	NominationID uint
}

// ParticipantManager struct
type ParticipantManager struct {
	db *DB
}

func NewParticipantManager(db *DB) (*ParticipantManager, error) {
	db.AutoMigrate(&Participant{})
	manager := ParticipantManager{}
	manager.db = db
	return &manager, nil
}

// HasParticipant -
func (state *ParticipantManager) HasParticipant(surname string) bool {
	if err := state.db.Where("surname=?", surname).Find(&Participant{}).Error; err != nil {
		return false
	}
	return true
}

// FindParticipant
func (state *ParticipantManager) FindParticipant(surname string) *Participant {
	participant := Participant{}
	state.db.Where("surname=?", surname).Find(&participant)
	return &participant
}

func (state *ParticipantManager) GetAll() []Participant {
	var participants []Participant
	state.db.Find(&participants)
	return participants
}

// NominationByID
func (state *ParticipantManager) ParticipantByID(id string) *Participant {
	participant := Participant{}
	state.db.First(&participant, id)
	return &participant
}

// DeleteParticipant
func (state *ParticipantManager) DeleteParticipant(id string) *Participant {
	participant := Participant{}
	state.db.Delete(&participant, id)
	return &participant
}

// CreateParticipant cre
func (state *ParticipantManager) CreateParticipant(surname string, name string, nomination *Nomination) *Participant {
	participant := Participant{
		Surname:    surname,
		Name:       name,
		Nomination: *nomination,
	}
	state.db.Create(&participant)
	return &participant
}
