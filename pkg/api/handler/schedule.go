package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func ScheduleHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	task := r.URL.Query().Get("task")
	if len(task) == 0 {
		messages := []string { "task not provided" }
		writeJsonResponse(w, nil, messages, http.StatusBadRequest)
		return
	}

	logger.Info("Scheduling task %s", task)

	switch task {
	case "alert-host":
		err := logic.AlertHost()
		if err != nil {
			logger.Error(err)
			messages := []string { err.Error() }
			writeJsonResponse(w, nil, messages, http.StatusInternalServerError)
		} else {
			writeJsonResponse(w, nil, nil, http.StatusOK)
		}
		return
	case "alert-guests":
		err := logic.AlertGuests()
		if err != nil {
			logger.Error(err)
			messages := []string { err.Error() }
			writeJsonResponse(w, nil, messages, http.StatusInternalServerError)
		} else {
			writeJsonResponse(w, nil, nil, http.StatusOK)
		}
		return
	default:
		err := fmt.Errorf("task %s not found", task)
		logger.Warn(err)
		messages := []string{ err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusNotFound)
		return
	}
}
