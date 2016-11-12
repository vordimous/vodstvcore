package dao

import (
	"esvodsApi/forms"
	"esvodsCore/models"

	"github.com/fatih/structs"
)

//VodDao ...
type VodDao struct{}

//Get ...
func (d VodDao) Get(id uint) (vod models.Vod, err error) {
	err = GetDB().First(&vod, id).Related(&vod.Tags, "Tags").Error
	return vod, err
}

//Find ...
func (d VodDao) Find(s forms.VodSearch) (vods []models.Vod, err error) {
	err = db.Where(getQuery(structs.New(s))).Find(&vods).Error
	return vods, err
}

//Save ...
func (d VodDao) Save(vod *models.Vod) (err error) {
	if GetDB().NewRecord(vod) {
		err = GetDB().Create(&vod).Error
	} else {
		err = GetDB().Save(&vod).Error
	}

	return err
}

//Delete ...
func (d VodDao) Delete(id uint) error {
	vod, err := d.Get(id)
	if err == nil {
		err = db.Delete(&vod).Error
	}
	return err
}
