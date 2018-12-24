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

// NominationManager struct
type NominationResultManager struct {
	db *DB
}

//NewNominationResultManager
func NewNominationResultManager(db *DB) (*NominationResultManager, error) {
	db.AutoMigrate(&NominationResult{})
	manager := NominationResultManager{}
	manager.db = db
	return &manager, nil
}

// ByNominationID
func (state *NominationResultManager) ByID(id string) *NominationResult {
	nominationResult := NominationResult{}
	state.db.First(&nominationResult, id)
	return &nominationResult
}

// ByNominationID
func (state *NominationResultManager) ByNominationID(nomination *Nomination) *NominationResult {
	nominationResult := NominationResult{}
	state.db.First(&nominationResult, "nomination_id = ?", nomination.ID)
	return &nominationResult
}

func (state *NominationResultManager) GetAll() []NominationResult {
	var results []NominationResult
	state.db.Find(&results)
	return results
}
