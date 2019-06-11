package model

import (
	"fmt"
)

type Mateo struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// NewMateoFromMap creates a Mateo from a database row
func NewMateoFromMap(row map[string]interface{}) (Mateo, error) {
	mateo := Mateo{}
	var ok bool

	if mateo.ID, ok = row["id"].(int64); !ok {
		return mateo, fmt.Errorf("field=id type=int64 not in row=%v", row)
	}
	if mateo.FirstName, ok = row["first_name"].(string); !ok {
		return mateo, fmt.Errorf("field=first_name type=string not in row=%v", row)
	}
	if mateo.LastName, ok = row["last_name"].(string); !ok {
		return mateo, fmt.Errorf("field=last_name type=string not in row=%v", row)
	}

	return mateo, nil
}