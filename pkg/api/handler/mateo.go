package handler

import (
	"bytes"
	htmlTemplate "html/template"
	"net/http"
	textTemplate "text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/email"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
	"github.com/wexel-nath/meat-night/pkg/model"
	"github.com/wexel-nath/meat-night/pkg/template"
)

func ListMateosHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	method := r.URL.Query().Get("method")
	mateos, err := logic.GetAllMateos(method)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, mateos, messages, http.StatusInternalServerError)
		return
	}

	if method == model.TypeLegacy {
		err = sendEmail(mateos[0])
		logger.LogIfErr(err)
	}

	writeJsonResponse(w, mateos, nil, http.StatusOK)
}

// TODO: move this to logic and change the trigger
func sendEmail(mateo model.Mateo) error {
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

	to := "nathanwelch_@hotmail.com" // mateo.Email
	message := email.Create(
		template.AlertHostSubject,
		textBuffer.String(),
		to,
	)
	message.SetHtml(htmlBuffer.String())

	return email.Send(message)
}
