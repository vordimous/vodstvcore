package dao

import "vodstv/core/models"

//TagDao ...
type TagDao struct{}

//Get ...
func (d TagDao) Get(id uint) (tag models.Tag, err error) {
	err = GetDB().First(&tag, id).Error
	return tag, err
}

//Find ...
func (d TagDao) Find(q map[string]interface{}) (tags []models.Tag, err error) {
	err = GetDB().Where(q).Find(&tags).Error
	return tags, err
}

//FindByTags ...
func (d TagDao) FindByTags(tagIDs []uint) (tags []models.Tag, err error) {
	if tagIDs != nil && len(tagIDs) > 0 {
		vtQuery := db.Table("vod_tags")
		vodIDs := []uint{}
		asstIDs := []uint{}
		err = vtQuery.Where("tag_id in (?)", tagIDs).Pluck("DISTINCT(vod_id)", &vodIDs).Error
		if err == nil {
			err = vtQuery.Where("vod_id in (?)", vodIDs).Pluck("DISTINCT(tag_id)", &asstIDs).Error
			if err == nil {
				err = GetDB().Where("id in (?)", asstIDs).Find(&tags).Error
			}
		}
	}
	return tags, err
}

//Save ...
func (d TagDao) Save(tag *models.Tag) (err error) {
	if GetDB().NewRecord(tag) {
		err = GetDB().Create(&tag).Error
	} else {
		err = GetDB().Model(&tag).Updates(getUpdates(tag)).Error
	}

	return err
}

//Delete ...
func (d TagDao) Delete(id uint) (tag models.Tag, err error) {
	tag, err = d.Get(id)
	if err == nil {
		err = GetDB().Delete(&tag).Error
	}
	return tag, err
}
