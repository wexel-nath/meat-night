package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/email"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func ScheduleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	task := r.URL.Query().Get("task")
	if len(task) == 0 {
		messages := []string { "task not provided" }
		writeJsonResponse(w, nil, messages, http.StatusBadRequest)
		return
	}

	logger.Info("Scheduling task %s", task)

	if task == "alert-host" {
		mateos, err := logic.GetAllMateos(model.TypeLegacy)
		if err != nil {
			logger.Error(err)
			messages := []string { err.Error() }
			writeJsonResponse(w, nil, messages, http.StatusInternalServerError)
			return
		}

		err = email.SendAlertHostEmail(mateos[0])
		if err != nil {
			logger.Error(err)
			messages := []string { err.Error() }
			writeJsonResponse(w, nil, messages, http.StatusInternalServerError)
			return
		}
		writeJsonResponse(w, nil, nil, http.StatusOK)
		return
	}
}
