package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Healthz(_ *http.Request, _ httprouter.Params) (interface{}, int, error) {
	result := struct{
		Status string `json:"status"`
	}{
		Status: "ok",
	}
	return result, http.StatusOK, nil
}
