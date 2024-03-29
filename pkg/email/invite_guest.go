package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	inviteGuestSubject = "Can you make it?"
)

func SendAlertGuestEmail(mateo model.Mateo, hostName string, inviteID string) error {
	html, text, err := createInviteGuestEmail(mateo.FirstName, hostName, inviteID)
	if err != nil {
		return err
	}

	message := create(inviteGuestSubject, html, text, mateo.Email)
	return send(message)
}

func createInviteGuestEmail(name string, hostName string, inviteID string) (string, string, error) {
	return build(hermes.Body{
		Name: name,
		Intros: []string{
			fmt.Sprintf("%s is up for meat night this week.", hostName),
		},
		Actions: []hermes.Action{
			{
				Instructions: "Can you make it?",
				Button: hermes.Button{
					Color:     "#22BC66",
					TextColor: "#ffffff",
					Text:      "I'm in!",
					Link:      buildAcceptInviteLink(model.TypeInviteGuest, inviteID),
				},
			},
			{
				Button: hermes.Button{
					Color:     "#EA4C25",
					TextColor: "#ffffff",
					Text:      "Busy making commish",
					Link:      buildDeclineInviteLink(model.TypeInviteGuest, inviteID),
				},
			},
		},
	})
}
