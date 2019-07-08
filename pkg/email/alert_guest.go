package email

import (
	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	AlertGuestSubject = "Meat Night"
)

func SendAlertGuestEmail(mateo model.Mateo) error {
	html, text, err := createAlertGuestEmail(mateo.FirstName)
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

func createAlertGuestEmail(name string) (string, string, error) {
	return build(hermes.Body{
		Name: name,
		Intros: []string{
			"<Mateo> is up for meat night this week.",
			"Let him know if you will attend.",
		},
	})
}
