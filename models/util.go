package models

//EsInt ...
type EsInt interface {
	FillStruct(m map[string]interface{}) error
}
