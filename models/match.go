package models

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

//Match ...
type Match struct {
	EsInt `json:"-"`
	gorm.Model
	Vods []Vod
}

//SetField ...
func (match *Match) SetField(name string, value interface{}) error {
	structValue := reflect.ValueOf(match).Elem()
	structFieldValue := structValue.FieldByName(name)

	return setField(&structValue, &structFieldValue, name, value)
}

// FillStruct ...
func (match *Match) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		if err := match.SetField(k, v); err != nil {
			return err
		}
	}
	return nil
}

//BeforeSave ...
// func (m Match) BeforeSave() (err error) {
//     var vs = []Match
//     dao.GetDB().Model(&m).Related(&vs)
//     var tags = make(map[string]bool)
//     for _, v := range vs) {
// 		var ts = []Match
//         dao.GetDB().Model(&v).Related(&ts)
//         for _, t := range ts) {
//             tags[t.Name] = true
//         }
// 	}
// 	return
// }
