package models

import (
	"reflect"
	"time"
)

//Tag ...
type Tag struct {
	EsInt     `json:"-"`
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
	Name      string     `json:"name"`
	Type      string     `json:"type"`
	Regex     string     `json:"regex"`
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
