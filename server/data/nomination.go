package data

import (
	"github.com/jinzhu/gorm"
	// import sqlite3 driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type Nomination struct {
	gorm.Model
	Name string
}

// UserManager struct
type NominationManager struct {
	db *DB
}

func NewNominationManager(db *DB) (*NominationManager, error) {
	db.AutoMigrate(&Nomination{})
	manager := NominationManager{}
	manager.db = db
	return &manager, nil
}

// HasParticipant -
func (state *NominationManager) HasNomination(name string) bool {
	if err := state.db.Where("name=?", name).Find(&Nomination{}).Error; err != nil {
		return false
	}
	return true
}

// FindParticipant
func (state *ParticipantManager) FindParticipant(name string) *Nomination {
	nomination := Nomination{}
	state.db.Where("name=?", name).Find(&nomination)
	return &nomination
}
