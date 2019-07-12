package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func Schedule(r *http.Request, _ httprouter.Params) (interface{}, int, error) {
	task := r.URL.Query().Get("task")
	if len(task) == 0 {
		return nil, http.StatusBadRequest, errors.New("task not provided")
	}

	logger.Info("Scheduling task %s", task)

	var err error

	switch task {
	case "alert-host":
		err = logic.InviteHost()
	case "guest-list":
		err = logic.GuestList()
	default:
		return nil, http.StatusNotFound, fmt.Errorf("task %s not found", task)
	}

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return nil, http.StatusOK, nil
}
