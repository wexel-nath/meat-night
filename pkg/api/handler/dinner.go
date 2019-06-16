package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
	"github.com/wexel-nath/meat-night/pkg/logic"
	"github.com/wexel-nath/meat-night/pkg/model"
)

func CreateDinnerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusBadRequest)
		return
	}

	var dinner model.Dinner
	err = json.Unmarshal(body, &dinner)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusBadRequest)
		return
	}

	newDinner, err := logic.CreateDinner(dinner)
	if err != nil {
		logger.Error(err)
		messages := []string { err.Error() }
		writeJsonResponse(w, nil, messages, http.StatusUnprocessableEntity)
		return
	}

	writeJsonResponse(w, newDinner, nil, http.StatusCreated)
}
