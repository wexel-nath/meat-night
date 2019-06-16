package model

import (
	"fmt"
	"time"
)

const (
	TypeLegacy = "legacy"
)

type Mateo struct {
	ID         int64   `json:"id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`

	// Fields for new sorting method
	GuestCount int64   `json:"guest_count"`
	HostCount  int64   `json:"host_count"`
	GuestRatio float64 `json:"guest_ratio"`

	// Legacy Fields
	LastHostDate string `json:"last_host_date"`
	Attended     int64  `json:"attended"`
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
	if mateo.GuestCount, ok = row["guest_count"].(int64); !ok {
		mateo.GuestCount = 0
	}
	if mateo.HostCount, ok = row["host_count"].(int64); !ok {
		mateo.HostCount = 0
	}
	if mateo.HostCount > 0 {
		mateo.GuestRatio = float64(mateo.GuestCount) / float64(mateo.HostCount)
	}
	if lastHostDate, ok := row["last_host_date"].(time.Time); ok {
		mateo.LastHostDate = lastHostDate.Format(DateFormat)
	}
	if mateo.Attended, ok = row["attended"].(int64); !ok {
		mateo.Attended = 0
	}

	return mateo, nil
}

var (
	mateoColumns = []string{
		"id",
		"first_name",
		"last_name",
		"guest_count",
		"host_count",
	}

	mateoLegacyColumns = []string{
		"mateo.id AS id",
		"first_name",
		"last_name",
		"last_host_date",
		"COUNT(*) AS attended",
	}
)

func GetMateoColumns() []string {
	return mateoColumns
}

func GetMateoLegacyColumns() []string {
	return mateoLegacyColumns
}

var (
	TestJohn = Mateo{
		ID:         1,
		FirstName:  "John",
		LastName:   "Doe",
		GuestCount: 5,
		HostCount:  2,
		GuestRatio: 2.5,
	}
	TestAdam = Mateo{
		ID:         2,
		FirstName:  "Adam",
		LastName:   "Samuel",
		GuestCount: 4,
		HostCount:  2,
		GuestRatio: 2.0,
	}

	// Legacy
	TestBob = Mateo{
		ID:           3,
		FirstName:    "Bob",
		LastName:     "Jane",
		LastHostDate: "12-05-19",
		Attended:     4,
	}
	TestDavid= Mateo{
		ID: 4,
		FirstName:    "David",
		LastName:     "Wilson",
		LastHostDate: "01-06-19",
		Attended:     1,
	}
)
