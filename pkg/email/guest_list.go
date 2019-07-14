package email

import (
	"fmt"
	"strings"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	guestListSubject = "Here's the guest list"
)

func SendGuestListEmail(host model.Mateo, dinnerID int64, invitees map[string][]string) error {
	html, text, err := createGuestListEmail(host.FirstName, dinnerID, invitees)
	if err != nil {
		return err
	}

	message := create(guestListSubject, html, text, host.Email)
	return send(message)
}

func createGuestListEmail(hostName string, dinnerID int64, invitees map[string][]string) (string, string, error) {
	return build(hermes.Body{
		Name: hostName,
		Intros: []string{
			"Here's who is coming to meat night this week",
		},
		Dictionary: buildGuestDictionary(invitees),
		Actions: []hermes.Action{
			{
				Instructions: "Where are you taking the lads?",
				Button: hermes.Button{
					Color:     "#1D89EB",
					TextColor: "#FFFFFF",
					Text:      "Click here to enter a venue",
					Link:      fmt.Sprintf("%s/dinner/%d/update", config.GetBaseURL(), dinnerID),
				},
			},
		},
	})
}

func buildGuestDictionary(invitees map[string][]string) []hermes.Entry {
	accepted := strings.Join(invitees["accepted"], ", ")
	declined := strings.Join(invitees["declined"], ", ")
	if declined == "" {
		declined= "-"
	}
	pending := strings.Join(invitees["pending"], ", ")
	if pending == "" {
		pending= "-"
	}

	return []hermes.Entry{
		{
			Key:   "Attending",
			Value: accepted,
		},
		{
			Key:   "Unavailable",
			Value: declined,
		},
		{
			Key:   "Waiting for response",
			Value: pending,
		},
	}
}
