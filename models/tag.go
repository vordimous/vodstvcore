package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

//Tag ...
type Tag struct {
	EsInt `json:"-"`
	gorm.Model
	Name  string
	Type  string
	Regex string
}

//SetField ...
func (tag *Tag) SetField(name string, value interface{}) error {
	structValue := reflect.ValueOf(tag).Elem()
	structFieldValue := structValue.FieldByName(name)

	return setField(&structValue, &structFieldValue, name, value)
}

// FillStruct ...
func (tag *Tag) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		if err := tag.SetField(k, v); err != nil {
			return err
		}
	}
	return nil
}
