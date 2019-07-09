package email

import (
	"fmt"
	"strings"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	AlertHostSubject = "Meat Night"
)

func SendAlertHostEmail(mateo model.Mateo, inviteID string) error {
	html, text, err := createAlertHostEmail(mateo.FirstName, inviteID)
	if err != nil {
		return err
	}

	message := create(AlertHostSubject, html, text, "nathanwelch_@hotmail.com")
	return send(message)
}

func createAlertHostEmail(name string, inviteID string) (string, string, error) {
	return build(hermes.Body{
		Name: name,
		Intros: []string{
			"You're up for meat night this week!",
		},
		Actions: []hermes.Action{
			{
				Instructions: "Let everyone know if you can make it:",
				Button: hermes.Button{
					Color: "#22BC66",
					Text:  "I'm available to host this week",
					Link:  buildInviteLink(model.TypeInviteHost, inviteID),
				},
			},
		},
	})
}

func buildInviteLink(inviteType string, inviteID string) string {
	return fmt.Sprintf("%s/invite/%s/%s", config.GetBaseURL(), strings.ToLower(inviteType), inviteID)
}
