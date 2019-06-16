package model

import (
	"fmt"
	"time"
)

const (
	TypeLegacy = "legacy"
)

var (
	mateoColumns = []string{
		"mateo_id",
		"first_name",
		"last_name",
		"total_attended",
		"total_hosted",
	}

	mateoLegacyColumns = []string{
		"mateo_id",
		"first_name",
		"last_name",
		"last_host_date",
		"COUNT(*) AS attended",
	}
)

type Mateo struct {
	ID         int64   `json:"mateo_id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`

	// Fields for new sorting method
	TotalAttended int64   `json:"total_attended"`
	TotalHosted   int64   `json:"total_hosted"`
	GuestRatio    float64 `json:"attendance_ratio"`

	// Legacy Fields
	LastHostDate string `json:"last_host_date"`
	Attended     int64  `json:"attended"`
}

// NewMateoFromMap creates a Mateo from a database row
func NewMateoFromMap(row map[string]interface{}) (Mateo, error) {
	mateo := Mateo{}
	var ok bool

	if mateo.ID, ok = row["mateo_id"].(int64); !ok {
		return mateo, fmt.Errorf("field=mateo_id type=int64 not in row=%v", row)
	}
	if mateo.FirstName, ok = row["first_name"].(string); !ok {
		return mateo, fmt.Errorf("field=first_name type=string not in row=%v", row)
	}
	if mateo.LastName, ok = row["last_name"].(string); !ok {
		return mateo, fmt.Errorf("field=last_name type=string not in row=%v", row)
	}
	if mateo.TotalAttended, ok = row["total_attended"].(int64); !ok {
		mateo.TotalAttended = 0
	}
	if mateo.TotalHosted, ok = row["total_hosted"].(int64); !ok {
		mateo.TotalHosted = 0
	}
	if mateo.TotalHosted > 0 {
		mateo.GuestRatio = float64(mateo.TotalAttended) / float64(mateo.TotalHosted)
	}
	if lastHostDate, ok := row["last_host_date"].(time.Time); ok {
		mateo.LastHostDate = lastHostDate.Format(DateFormat)
	}
	if mateo.Attended, ok = row["attended"].(int64); !ok {
		mateo.Attended = 0
	}

	return mateo, nil
}

func GetMateoColumns() []string {
	return mateoColumns
}

func GetMateoLegacyColumns() []string {
	return mateoLegacyColumns
}
