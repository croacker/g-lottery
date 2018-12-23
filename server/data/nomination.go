package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Nomination struct
type Nomination struct {
	gorm.Model
	Name string
}

// NominationManager struct
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

// FindNomination
func (state *NominationManager) NominationByName(name string) *Nomination {
	nomination := Nomination{}
	state.db.Where("name=?", name).Find(&nomination)
	return &nomination
}

// NominationByID
func (state *NominationManager) NominationByID(id string) *Nomination {
	nomination := Nomination{}
	state.db.First(&nomination, id)
	return &nomination
}

// DeleteNomination
func (state *NominationManager) DeleteNomination(id string) *Nomination {
	nomination := Nomination{}
	state.db.Delete(&nomination, id)
	return &nomination
}

// CreateNomination
func (state *NominationManager) CreateNomination(name string) *Nomination {
	nomination := Nomination{
		Name: name,
	}
	state.db.Create(&nomination)
	return &nomination
}

// GetAll ...
func (state *NominationManager) GetAll() []Nomination {
	var nominations []Nomination
	state.db.Find(&nominations)
	return nominations
}

//Predefined
func (state *NominationManager) Predefined() {
	name := "The best hunter"
	if !state.HasNomination(name) {
		state.CreateNomination(name)
	}

	name = "Креативный класс"
	if !state.HasNomination(name) {
		state.CreateNomination(name)
	}

	name = "Наставник года"
	if !state.HasNomination(name) {
		state.CreateNomination(name)
	}

	name = "Сертификаты"
	if !state.HasNomination(name) {
		state.CreateNomination(name)
	}

	name = "Спасибо"
	if !state.HasNomination(name) {
		state.CreateNomination(name)
	}

	name = "Участник трансформации"
	if !state.HasNomination(name) {
		state.CreateNomination(name)
	}
}
