package email

import (
	"fmt"

	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	AlertGuestSubject = "Meat Night"
)

func SendAlertGuestEmail(mateo model.Mateo, hostName string) error {
	html, text, err := createAlertGuestEmail(mateo.FirstName, hostName)
	if err != nil {
		return err
	}

	message := create(
		AlertGuestSubject,
		html,
		text,
		"nathanwelch_@hotmail.com",
		//mateo.Email,
	)

	return send(message)
}

func createAlertGuestEmail(name string, hostName string) (string, string, error) {
	return build(hermes.Body{
		Name: name,
		Intros: []string{
			fmt.Sprintf("%s is up for meat night this week.", hostName),
			"Let him know if you will attend.",
		},
	})
}
