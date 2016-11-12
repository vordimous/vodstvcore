package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

//Vod ...
type Vod struct {
	EsInt `json:"-"`
	gorm.Model
	MatchID   uint
	Tags      []Tag `gorm:"many2many:vod_tags;"`
	Title     string
	Content   string
	VideoKey  string
	VideoURL  string
	VideoSrc  string
	VideoDate string
	ThumbURL  string
}

//SetField ...
func (v *Vod) SetField(name string, value interface{}) error {
	sv := reflect.ValueOf(v).Elem()
	sf := sv.FieldByName(name)

	return setField(&sv, &sf, name, value)
}

// FillStruct ...
func (v *Vod) FillStruct(m map[string]interface{}) error {
	for k, val := range m {
		if err := v.SetField(k, val); err != nil {
			return err
		}
	}
	return nil
}
