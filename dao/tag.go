package dao

import (
	"esvodsApi/forms"
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//TagDao ...
type TagDao struct{}

//Get ...
func (d TagDao) Get(id uint) (tag models.Tag, err error) {
	err = GetDB().First(&tag, id).Error
	return tag, err
}

//Find ...
func (d TagDao) Find(s forms.TagSearch) (tags []models.Tag, err error) {
	err = db.Where(getQuery(structs.New(s))).Find(&tags).Error
	return tags, err
}

//Save ...
func (d TagDao) Save(tag *models.Tag) (err error) {
	if GetDB().NewRecord(tag) {
		err = GetDB().Create(&tag).Error
	} else {
		err = GetDB().Save(&tag).Error
	}

	return err
}

//Delete ...
func (d TagDao) Delete(id uint) error {
	tag, err := d.Get(id)
	if err == nil {
		err = db.Delete(&tag).Error
	}
	return err
}
