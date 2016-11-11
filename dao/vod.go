package dao

import (
	"esvodsApi/forms"
	"esvodsCore/models"
)

//VodDao ...
type VodDao struct{}

//Get ...
func (d VodDao) Get(id uint) (vod models.Vod, err error) {
	err = GetDB().First(&vod, id).Error
	return vod, err
}

//Find ...
func (d VodDao) Find(vodSearch forms.VodSearch) (vods []models.Vod, err error) {
	// _, err = db.GetDB().Select(&vods, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS watcher FROM vod a LEFT JOIN public.watcher u ON a.watcher_id = u.id WHERE a.watcher_id=$1 GROUP BY a.id, a.title, a.content, a.updated_at, a.created_at, u.id, u.name, u.email ORDER BY a.id DESC", watcherID)
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
func (d VodDao) Delete(id uint) (err error) {
	// _, err = d.One(watcherID, id)

	// if err != nil {
	// 	return errors.New("models.Vod not found")
	// }

	// _, err = db.GetDB().Exec("DELETE FROM vod WHERE id=$1", id)

	return err
}
