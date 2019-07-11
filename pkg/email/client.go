package email

import (
	"fmt"

	"github.com/mailgun/mailgun-go"
	"github.com/matcornic/hermes/v2"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/logger"
)

var(
	builder hermes.Hermes
	client  mailgun.Mailgun
)

func Configure() {
	builder = hermes.Hermes{
		Product: hermes.Product{
			Name: config.GetCompanyName(),
			Logo: "https://i.ibb.co/8PTB7ZF/mateo-corp-transparent.png",
			Copyright: "Copyright © 2019 Wexel Tech. All rights reserved.",
		},
	}

	domain := config.GetMailgunDomain()
	apiKey := config.GetMailgunApiKey()

	client = mailgun.NewMailgun(domain, apiKey)
}

func build(body hermes.Body) (string, string, error) {
	email := hermes.Email{ Body: body }

	html, err := builder.GenerateHTML(email)
	if err != nil {
		return "", "", err
	}

	text, err := builder.GeneratePlainText(email)
	if err != nil {
		return "", "", err
	}

	return html, text, nil
}

func create(subject string, html string, text string, to ...string) *mailgun.Message {
	from := fmt.Sprintf("%s <%s>", config.GetCompanyName(), config.GetCompanyEmail())
	message := client.NewMessage(from, subject, text, to...)
	message.SetHtml(html)
	return message
}

func send(message *mailgun.Message) error {
	resp, id, err := client.Send(message)
	logger.Info("Email sent to Mailgun. resp=%s id=%s err=%v", resp, id, err)
	return err
}
