package email

import (
	"fmt"

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

func create(subject string, text string, to ...string) *mailgun.Message {
	from := fmt.Sprintf("%s <%s>", config.GetCompanyName(), config.GetCompanyEmail())
	return client.NewMessage(from, subject, text, to...)
}

func send(message *mailgun.Message) error {
	resp, id, err := client.Send(message)
	logger.Info("Email sent to Mailgun. resp=%s id=%s err=%v", resp, id, err)
	return err
}
