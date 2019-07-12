package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/logger"
)

type response struct {
	Result   interface{} `json:"result"`
	Messages []string    `json:"messages"`
}

func newResponse(result interface{}, messages []string) response {
	return response{
		Result:   result,
		Messages: messages,
	}
}

func writeJsonResponse(
	resp http.ResponseWriter,
	result interface{},
	messages []string,
	status int,
) {
	bytes, err := json.Marshal(newResponse(result, messages))
	if err != nil {
		logger.Error(err)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	resp.Write(bytes)
}

type handleFunc func(r *http.Request, ps httprouter.Params) (interface{}, int, error)

func requestHandler(handler handleFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		messages := make([]string, 0)

		result, statusCode, err := handler(r, ps)
		if err != nil {
			logger.Error(err)
			messages = []string{ err.Error() }
		}

		writeJsonResponse(w, result, messages, statusCode)
	}
}
