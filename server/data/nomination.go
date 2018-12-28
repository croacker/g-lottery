package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Nomination struct
type Nomination struct {
	gorm.Model
	Name string
	Code string
}

// NominationManager struct
type NominationManager struct {
	db *DB
}

// NewNominationManager factory method
func NewNominationManager(db *DB) (*NominationManager, error) {
	db.AutoMigrate(&Nomination{})
	manager := NominationManager{}
	manager.db = db
	return &manager, nil
}

// HasNomination - check Nomination by code
func (state *NominationManager) HasNomination(code string) bool {
	if err := state.db.Where("code=?", code).Find(&Nomination{}).Error; err != nil {
		return false
	}
	return true
}

// NominationByName get Nomination by name
func (state *NominationManager) NominationByName(name string) *Nomination {
	nomination := Nomination{}
	state.db.Where("name=?", name).Find(&nomination)
	return &nomination
}

// NominationByCode get Nomination by code
func (state *NominationManager) NominationByCode(code string) *Nomination {
	nomination := Nomination{}
	state.db.Where("code=?", code).Find(&nomination)
	return &nomination
}

// NominationByID get Nomination by id
func (state *NominationManager) NominationByID(id string) *Nomination {
	nomination := Nomination{}
	state.db.First(&nomination, id)
	return &nomination
}

// DeleteNomination delete Nomination by id
func (state *NominationManager) DeleteNomination(id string) *Nomination {
	nomination := Nomination{}
	state.db.Delete(&nomination, id)
	return &nomination
}

// CreateNomination create Nomination by name and code
func (state *NominationManager) CreateNomination(name string, code string) *Nomination {
	nomination := Nomination{
		Name: name,
		Code: code,
	}
	state.db.Create(&nomination)
	return &nomination
}

// GetAll get All Nominations
func (state *NominationManager) GetAll() []Nomination {
	var nominations []Nomination
	state.db.Find(&nominations)
	return nominations
}

//Predefined insert predefined Nominations
func (state *NominationManager) Predefined() {
	name := "The best hunter"
	code := "best-hunter"
	if !state.HasNomination(code) {
		state.CreateNomination(name, code)
	}

	name = "Креативный класс"
	code = "creative-class"
	if !state.HasNomination(code) {
		state.CreateNomination(name, code)
	}

	name = "Наставник года"
	code = "menthor"
	if !state.HasNomination(code) {
		state.CreateNomination(name, code)
	}

	name = "Сертификаты"
	code = "certificate"
	if !state.HasNomination(code) {
		state.CreateNomination(name, code)
	}

	name = "Спасибо"
	code = "thanks"
	if !state.HasNomination(code) {
		state.CreateNomination(name, code)
	}

	name = "Участник трансформации"
	code = "transformation"
	if !state.HasNomination(code) {
		state.CreateNomination(name, code)
	}
}
