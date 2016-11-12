package dao

import (
	"esvodsApi/forms"
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//MatchDao ...
type MatchDao struct{}

//Get ...
func (d MatchDao) Get(id uint) (match models.Match, err error) {
	err = GetDB().First(&match, id).Error
	return match, err
}

//Find ...
func (d MatchDao) Find(s forms.MatchSearch) (matchs []models.Match, err error) {
	err = db.Where(getQuery(structs.New(s))).Find(&matchs).Error
	return matchs, err
}

//Save ...
func (d MatchDao) Save(match *models.Match) (err error) {
	if GetDB().NewRecord(match) {
		err = GetDB().Create(&match).Error
	} else {
		err = GetDB().Save(&match).Error
	}

	return err
}

//Delete ...
func (d MatchDao) Delete(id uint) error {
	match, err := d.Get(id)
	if err == nil {
		err = db.Delete(&match).Error
	}
	return err
}
