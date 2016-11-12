package dao

import (
	"strings"

	"github.com/fatih/structs"
)

func getQuery(s *structs.Struct) map[string]interface{} {
	q := make(map[string]interface{})
	for _, f := range s.Fields() {
		q[strings.ToLower(f.Name())] = f.Value()
	}
	return q
}
