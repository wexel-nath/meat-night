package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	guestListSubject = "Here's the guest list"
)

func SendGuestListEmail(host model.Mateo, guestNames []string) error {
	html, text, err := createGuestListEmail(host.FirstName, guestNames)
	if err != nil {
		return err
	}

	message := create(guestListSubject, html, text, "nathanwelch_@hotmail.com")
	return send(message)
}

func createGuestListEmail(hostName string, guestNames []string) (string, string, error) {
	entries := buildEntries(hostName, guestNames)
	return build(hermes.Body{
		Name: hostName,
		Intros: []string{
			fmt.Sprintf("These %d legends will be attending", len(entries)),
		},
		Table: hermes.Table{
			Data: entries,
		},
	})
}

func buildEntries(hostName string, guestNames []string) [][]hermes.Entry {
	entries := make([][]hermes.Entry, 0)
	for _, name := range guestNames {
		if name == hostName {
			name = "You"
		}
		entries = append(entries, []hermes.Entry{
			{ Value: name },
		})
	}

	return entries
}

// TODO:
// triggered via a cron && email has not been sent before
// X people have accepted, Y people have declined, Z people haven't replied

// triggered once everyone has replied
// X people have accepted, Y people have declined
//  => after each guest invite response
//  => check pending invites for dinner id
