package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NominationResult struct
type NominationResult struct {
	gorm.Model
	Participant   Participant `gorm:"foreignkey:ParticipantId"`
	ParticipantID uint
	Nomination    Nomination `gorm:"foreignkey:NominationId"`
	NominationID  uint
}

// NominationResultManager struct
type NominationResultManager struct {
	db *DB
}

//NewNominationResultManager factory method
func NewNominationResultManager(db *DB) (*NominationResultManager, error) {
	db.AutoMigrate(&NominationResult{})
	manager := NominationResultManager{}
	manager.db = db
	return &manager, nil
}

// ByID get NominationResult by id
func (state *NominationResultManager) ByID(id string) *NominationResult {
	nominationResult := NominationResult{}
	state.db.First(&nominationResult, id)
	return &nominationResult
}

// ByNominationID get NominationResult for Nomination
func (state *NominationResultManager) ByNominationID(nomination *Nomination) *NominationResult {
	nominationResult := NominationResult{}
	state.db.Preload("Participant").First(&nominationResult, "nomination_id = ?", nomination.ID)
	return &nominationResult
}

// CreateNominationResult create NominationResult for Nomination and Participant
func (state *NominationResultManager) CreateNominationResult(nomination *Nomination, participant *Participant) *NominationResult {
	nominationResult := NominationResult{
		ParticipantID: participant.ID,
		NominationID:  nomination.ID,
	}
	state.db.Create(&nominationResult)
	return &nominationResult
}

// GetAll get All NominationResults
func (state *NominationResultManager) GetAll() []NominationResult {
	var results []NominationResult
	state.db.Preload("Participant").Preload("Nomination").Find(&results)
	return results
}

// Delete delete NominationResult
func (state *NominationResultManager) Delete(nominationResult *NominationResult) {
	state.db.Delete(&nominationResult)
}
