package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logic"
)

func ListMateos(r *http.Request, ps httprouter.Params) (interface{}, int, error) {
	method := r.URL.Query().Get("method")

	mateos, err := logic.GetAllMateos(method)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return mateos, http.StatusOK, nil
}
