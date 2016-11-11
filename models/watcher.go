package models

import (
	"esvodsCore/util"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//Watcher ...
type Watcher struct {
	EsInt
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}

//BeforeCreate ...
func (w *Watcher) BeforeCreate() (err error) {
	bytePassword := []byte(w.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	util.CheckErr(err, "Pass hash failed")

	w.Password = string(hashedPassword)
	return
}
