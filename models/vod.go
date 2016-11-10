package models

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jinzhu/gorm"
)

//Vod ...
type Vod struct {
	EsInt `json:"_"`
	gorm.Model
	Title   string
	Content string
}

func setField(structValue *reflect.Value, structFieldValue *reflect.Value, name string, value interface{}) error {
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

//SetField ...
func (vod *Vod) SetField(name string, value interface{}) error {
	structValue := reflect.ValueOf(vod).Elem()
	structFieldValue := structValue.FieldByName(name)

	return setField(&structValue, &structFieldValue, name, value)
}

// FillStruct ...
func (vod *Vod) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := vod.SetField(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
