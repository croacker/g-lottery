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
	Department   string
	Chance       int
	Nomination   Nomination `gorm:"foreignkey:NominationId"`
	NominationID uint
}

// ParticipantManager struct
type ParticipantManager struct {
	db *DB
}

// NewParticipantManager factory method
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

// GetAll get all Participants
func (state *ParticipantManager) GetAll() []Participant {
	var participants []Participant
	state.db.Find(&participants)
	return participants
}

// ByID get Participant by id
func (state *ParticipantManager) ByID(id string) *Participant {
	participant := Participant{}
	state.db.First(&participant, id)
	return &participant
}

// ByNominationID Participant by Nomination
func (state *ParticipantManager) ByNominationID(nomination *Nomination) []Participant {
	participants := []Participant{}
	state.db.Where("nomination_id = ?", nomination.ID).Find(&participants)
	return participants
}

// DeleteById delete Participant by id
func (state *ParticipantManager) DeleteById(id string) *Participant {
	participant := Participant{}
	state.db.Delete(&participant, id)
	return &participant
}

// Delete delete concrete Participant
func (state *ParticipantManager) Delete(participant *Participant) {
	state.db.Delete(&participant)
}

// CreateParticipant create Participant by surname, name, department, nomination
func (state *ParticipantManager) CreateParticipant(surname string, name string, department string, chance int, nomination *Nomination) *Participant {
	participant := Participant{
		Surname:      surname,
		Name:         name,
		Department:   department,
		Chance:       chance,
		NominationID: nomination.ID,
	}
	state.db.Create(&participant)
	return &participant
}
