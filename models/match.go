package models

import "time"

//Match ...
type Match struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"_"`
	UpdatedAt time.Time  `json:"_"`
	DeletedAt *time.Time `json:"_" sql:"index"`
	Title     string     `json:"title"`
	Vods      []Vod      `json:"vods"`
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
