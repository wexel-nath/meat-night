package email

import (
	"github.com/mailgun/mailgun-go"
	"github.com/wexel-nath/meat-night/pkg/config"
	"github.com/wexel-nath/meat-night/pkg/logger"
)

var(
	client mailgun.Mailgun
)

func ConfigureClient() {
	domain := config.GetMailgunDomain()
	apiKey := config.GetMailgunApiKey()

	client = mailgun.NewMailgun(domain, apiKey)
}

func Create(from, subject, text string, to ...string) *mailgun.Message {
	return client.NewMessage(from, subject, text, to...)
}

func Send(message *mailgun.Message) error {
	resp, id, err := client.Send(message)
	logger.Info("Email sent to Mailgun. resp=%s id=%s err=%s", resp, id, err.Error())
	return err
}
