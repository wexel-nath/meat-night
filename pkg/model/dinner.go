package model

import "time"

const (
	DateFormat = "02-01-06"
)

type Dinner struct {
	ID     int64     `json:"id"`
	Date   time.Time `json:"date"`
	Host   Mateo     `json:"host"`
	Venue  string    `json:"venue"`
	Guests []Mateo   `json:"guests"`
}

type DinnerRequestDto struct {
	Date     string  `json:"date"`
	HostID   int64   `json:"host_id"`
	Venue    string  `json:"venue"`
	GuestIDs []int64 `json:"guest_ids"`
}
