package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/email"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
	"github.com/wexel-nath/meat-night/pkg/model"
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
		sendEmail(mateos[0])
	}

	writeJsonResponse(w, mateos, nil, http.StatusOK)
}

func sendEmail(mateo model.Mateo) {
	subject := "Meat Night"
	bodyFormat := `
Mateo,
This week is %s's turn.
Last time he took the mateos to <here>.
Will this be a top 5?
	`
	body := fmt.Sprintf(bodyFormat, mateo.FirstName)
	to := "nathanwelch_@hotmail.com"

	message := email.Create(
		subject,
		body,
		to,
	)
	err := email.Send(message)
	logger.LogIfErr(err)
}
