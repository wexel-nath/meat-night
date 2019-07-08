package email

import (
	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/model"
)

const (
	AlertHostSubject = "Meat Night"
)

func SendAlertHostEmail(mateo model.Mateo) error {
	html, text, err := createAlertHostEmail(mateo.FirstName)
	if err != nil {
		return err
	}

	message := create(
		AlertHostSubject,
		html,
		text,
		"nathanwelch_@hotmail.com",
		//mateo.Email,
	)

	return send(message)
}

func createAlertHostEmail(name string) (string, string, error) {
	return build(hermes.Body{
		Name: name,
		Intros: []string{
			"You're up for meat night this week!",
			"Let everyone know if you can make it or not.",
		},
	})
}
