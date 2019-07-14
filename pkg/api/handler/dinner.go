package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/config"
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

func UpdateVenue(r *http.Request, _ httprouter.Params) (string, error) {
	return "thumbs-up", updateVenue(r)
}

func updateVenue(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	var venue string
	var dinnerID int64

	formData := strings.Split(string(body), "&")
	for _, values := range formData {
		v := strings.Split(values, "=")
		if len(v) != 2 {
			return fmt.Errorf("bad form values: %v", string(body))
		}

		key := v[0]
		value := v[1]

		switch key {
		case "venue":
			venue = strings.ReplaceAll(value, "+", " ")
		case "dinner-id":
			dinnerID, err = strconv.ParseInt(value, 10, 64)
			if err != nil {
				return err
			}
		}
	}

	return logic.UpdateDinnerVenue(dinnerID, venue)
}

func UpdateVenueForm(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	dinnerID := ps.ByName("dinnerID")
	w.Write([]byte(getUpdateVenueForm(dinnerID)))
}

func getUpdateVenueForm(dinnerID string) string {
	actionUrl := config.GetBaseURL() + "/dinner/update"
	form := `
		<!DOCTYPE html>
		<html>
			<head>
				<title>Meat Night</title>
				<link rel="shortcut icon" href="http://www.iconj.com/ico/8/9/89rzu5yyr7.ico" type="image/x-icon" />
			</head>
			<body>
				<div align="center">
					<form action="%s" method="POST">
						<img src="%s" height="200px" width="200px" />
						<h3>Where are you taking the lads?</h3>
						<input type="text" name="venue" />
						<input type="hidden" name="dinner-id" value="%s" />
						<button type="submit">Submit</button>
					</form>
				</div>
			</body>
		</html>
	`
	return fmt.Sprintf(form, actionUrl, config.GetCompanyLogo(), dinnerID)
}
