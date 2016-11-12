package models

import (
	"errors"
	"fmt"
	"reflect"
)

//EsInt ...
type EsInt interface {
	FillStruct(m map[string]interface{}) error
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
