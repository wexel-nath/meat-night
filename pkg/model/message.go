package model

import (
	"fmt"
	"time"
)

const (
	TypeAlertHost = "alert.host"
	TypeAlertGuest = "alert.guest"
)

var (
	messageColumns = []string{
		"message_id",
		"message_event",
		"message_timestamp",
		"mateo_id",
	}
)

type Message struct {
	ID        int64  `json:"message_id"`
	Event     string `json:"message_event"`
	Timestamp string `json:"message_timestamp"`
	MateoID   int64  `json:"mateo_id"`
}

func NewMessageFromMap(row map[string]interface{}) (Message, error) {
	message := Message{}
	var ok bool

	if message.ID, ok = row["message_id"].(int64); !ok {
		return message, fmt.Errorf("field=message_id type=int64 not in row=%v", row)
	}
	if message.Event, ok = row["message_event"].(string); !ok {
		return message, fmt.Errorf("field=message_event type=string not in row=%v", row)
	}
	if timestamp, ok := row["message_timestamp"].(time.Time); ok {
		message.Timestamp = timestamp.Format(time.RFC1123)
	}
	if message.MateoID, ok = row["mateo_id"].(int64); !ok {
		return message, fmt.Errorf("field=mateo_id type=int64 not in row=%v", row)
	}

	return message, nil
}

func GetMessageColumns() []string {
	return messageColumns
}
