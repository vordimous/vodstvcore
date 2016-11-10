package models

import "github.com/jinzhu/gorm"

//Watcher ...
type Watcher struct {
	EsInt
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
}
