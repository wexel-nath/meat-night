package model

import "fmt"

const (
	// invite status types
	TypeInvitePending  = "PENDING"
	TypeInviteAccepted = "ACCEPTED"
	TypeInviteDeclined = "DECLINED"

	// invite types
	TypeInviteHost  = "HOST"
	TypeInviteGuest = "GUEST"
)

var (
	inviteColumns = []string{
		"invite_id",
		"invite_type",
		"invite_status",
		"mateo_id",
		"dinner_id",
	}
)

func GetInviteColumns() []string {
	return inviteColumns
}

type Invite struct {
	InviteID     string `json:"invite_id"`
	InviteType   string `json:"invite_type"`
	InviteStatus string `json:"invite_status"`
	MateoID      int64  `json:"mateo_id"`
	DinnerID     *int64  `json:"dinner_id"`
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
	if invite.DinnerID, ok = row["dinner_id"].(*int64); !ok {
		return invite, fmt.Errorf("field=dinner_id type=int64 not in row=%v", row)
	}

	return invite, nil
}
