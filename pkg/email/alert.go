package email

import (
	"bytes"
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/wexel-nath/meat-night/pkg/email/template"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func SendAlertHostEmail(mateo model.Mateo) error {
	alertHostTextTemplate, err := textTemplate.New("AlertHostText").Parse(template.AlertHostText)
	if err != nil {
		return err
	}

	var textBuffer bytes.Buffer
	if err = alertHostTextTemplate.Execute(&textBuffer, mateo); err != nil {
		return err
	}

	alertHostHtmlTemplate, err := htmlTemplate.New("AlertHostHtml").Parse(template.AlertHostHtml)
	if err != nil {
		return err
	}

	var htmlBuffer bytes.Buffer
	if err = alertHostHtmlTemplate.Execute(&htmlBuffer, mateo); err != nil {
		return err
	}

	message := create(
		template.AlertHostSubject,
		textBuffer.String(),
		"nathanwelch_@hotmail.com",
		mateo.Email,
	)
	message.SetHtml(htmlBuffer.String())

	return send(message)
}
