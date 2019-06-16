package model

import (
	"fmt"
	"time"
)

const (
	DateFormat = "02-01-06"
)

var (
	dinnerColumns = []string{
		"dinner_id",
		"date",
		"venue",
		"last_name",
	}
)

type Dinner struct {
	ID       int64    `json:"dinner_id"`
	Date     string   `json:"date"`
	Venue    string   `json:"venue"`
	Host     string   `json:"host"`
	Attended []string `json:"attended"`
}

func NewDinnerFromMap(row map[string]interface{}) (Dinner, error) {
	dinner := Dinner{}
	var ok bool

	if dinner.ID, ok = row["dinner_id"].(int64); !ok {
		return dinner, fmt.Errorf("field=dinner_id type=int64 not in row=%v", row)
	}
	if date, ok := row["date"].(time.Time); ok {
		dinner.Date = date.Format(DateFormat)
	}
	if dinner.Venue, ok = row["venue"].(string); !ok {
		return dinner, fmt.Errorf("field=venue type=string not in row=%v", row)
	}
	if dinner.Host, ok = row["last_name"].(string); !ok {
		return dinner, fmt.Errorf("field=last_name type=string not in row=%v", row)
	}

	return dinner, nil
}

type DinnerRequestDto struct {
	Date     string  `json:"date"`
	HostID   int64   `json:"host_id"`
	Venue    string  `json:"venue"`
	GuestIDs []int64 `json:"guest_ids"`
}

func GetDinnerColumns() []string {
	return dinnerColumns
}
