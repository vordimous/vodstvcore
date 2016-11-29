package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

//Feed ...
type Feed struct {
	EsInt `json:"-"`
	gorm.Model
	OwnerID uint
	Name    string
	Tags    []Tag `gorm:"many2many:feed_tags;"`
}

//SetField ...
func (feed *Feed) SetField(name string, value interface{}) error {
	structValue := reflect.ValueOf(feed).Elem()
	structFieldValue := structValue.FieldByName(name)

	return setField(&structValue, &structFieldValue, name, value)
}

// FillStruct ...
func (feed *Feed) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		if err := feed.SetField(k, v); err != nil {
			return err
		}
	}
	return nil
}
