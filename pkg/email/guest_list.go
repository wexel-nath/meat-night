package email

import (
	"strings"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	guestListSubject = "Here's the guest list"
)

func SendGuestListEmail(host model.Mateo, invitees map[string][]string) error {
	html, text, err := createGuestListEmail(host.FirstName, invitees)
	if err != nil {
		return err
	}

	message := create(guestListSubject, html, text, host.Email)
	return send(message)
}

func createGuestListEmail(hostName string, invitees map[string][]string) (string, string, error) {
	return build(hermes.Body{
		Name: hostName,
		Intros: []string{
			"Here's who is coming to meat night this week",
		},
		Dictionary: buildGuestDictionary(invitees),
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

// TODO:
// triggered via a cron && email has not been sent before
// X people have accepted, Y people have declined, Z people haven't replied

// triggered once everyone has replied
// X people have accepted, Y people have declined
//  => after each guest invite response
//  => check pending invites for dinner id
