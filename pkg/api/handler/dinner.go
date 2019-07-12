package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func CreateDinner(r *http.Request, _ httprouter.Params) (interface{}, int, error) {
	// todo: requires auth

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	var request struct {
		Date     string   `json:"date"`
		Venue    string   `json:"venue"`
		Host     string   `json:"host"`
		Attended []string `json:"attended"`
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	dinner, err := logic.CreateDinner(request.Date, request.Venue, request.Host, request.Attended)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	return dinner, http.StatusCreated, nil
}
