package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	inviteHostSubject = "You're up!"
)

func SendAlertHostEmail(mateo model.Mateo, inviteID string) error {
	html, text, err := createInviteHostEmail(mateo.FirstName, inviteID)
	if err != nil {
		return err
	}

	message := create(inviteHostSubject, html, text, mateo.Email)
	return send(message)
}

func createInviteHostEmail(name string, inviteID string) (string, string, error) {
	return build(hermes.Body{
		Name: name,
		Intros: []string{
			"You're up for meat night this week!",
		},
		Actions: []hermes.Action{
			{
				Instructions: "Let everyone know if you can make it:",
				Button: hermes.Button{
					Color:     "#22BC66",
					TextColor: "#ffffff",
					Text:      "I'm available to host",
					Link:      buildAcceptInviteLink(model.TypeInviteHost, inviteID),
				},
			},
			{
				Button: hermes.Button{
					Color:     "#EA4C25",
					TextColor: "#ffffff",
					Text:      "I'm up the coast",
					Link:      buildDeclineInviteLink(model.TypeInviteHost, inviteID),
				},
			},
		},
	})
}

func buildAcceptInviteLink(inviteType string, inviteID string) string {
	return fmt.Sprintf("%s/%s/%s/accept", config.GetBaseURL(), inviteType, inviteID)
}

func buildDeclineInviteLink(inviteType string, inviteID string) string {
	return fmt.Sprintf("%s/%s/%s/decline", config.GetBaseURL(), inviteType, inviteID)
}
