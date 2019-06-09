package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
)

func HealthzHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		logger.Error(err)
	}
}
