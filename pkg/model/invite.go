package model

import (
	"fmt"
	"time"
)

const (
	// invite status types
	TypeInvitePending  = "pending"
	TypeInviteAccepted = "accepted"
	TypeInviteDeclined = "declined"

	// invite types
	TypeInviteHost  = "host"
	TypeInviteGuest = "guest"
)

var (
	inviteColumns = []string{
		"invite_id",
		"invite_type",
		"invite_status",
		"mateo_id",
		"dinner_id",
		"dinner_time",
		"timestamp",
	}
)

func GetInviteColumns() []string {
	return inviteColumns
}

type Invite struct {
	InviteID     string    `json:"invite_id"`
	InviteType   string    `json:"invite_type"`
	InviteStatus string    `json:"invite_status"`
	MateoID      int64     `json:"mateo_id"`
	DinnerID     int64     `json:"dinner_id"`
	DinnerTime   time.Time `json:"dinner_time"`
	Timestamp    time.Time `json:"timestamp"`
}

func NewInviteFromRow(row map[string]interface{}) (Invite, error) {
	invite := Invite{}
	var ok bool

	if invite.InviteID, ok = row["invite_id"].(string); !ok {
		return invite, fmt.Errorf("field=invite_id type=string not in row=%v", row)
	}
	if invite.InviteType, ok = row["invite_type"].(string); !ok {
		return invite, fmt.Errorf("field=invite_type type=string not in row=%v", row)
	}
	if invite.InviteStatus, ok = row["invite_status"].(string); !ok {
		return invite, fmt.Errorf("field=invite_status type=string not in row=%v", row)
	}
	if invite.MateoID, ok = row["mateo_id"].(int64); !ok {
		return invite, fmt.Errorf("field=mateo_id type=int64 not in row=%v", row)
	}
	if invite.DinnerID, ok = row["dinner_id"].(int64); !ok {
		invite.DinnerID = 0
	}
	if invite.DinnerTime, ok = row["dinner_time"].(time.Time); !ok {
		invite.DinnerTime = time.Time{}
	}
	if invite.Timestamp, ok = row["timestamp"].(time.Time); !ok {
		return invite, fmt.Errorf("field=timestamp type=time.Time not in row=%v", row)
	}

	return invite, nil
}
