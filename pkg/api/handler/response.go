package handler

import (
	"encoding/json"
	"net/http"

	"github.com/wexel-nath/meat-night/pkg/logger"
)

type message struct {
	Code    string `json:"code"`
	Level   string `json:"level"`
	Message string `json:"message"`
}

type response struct {
	Result   interface{} `json:"result"`
	Messages []message   `json:"messages"`
}

func newResponse(result interface{}, messages []message) response {
	return response{
		Result:   result,
		Messages: messages,
	}
}

func writeJsonResponse(
	resp http.ResponseWriter,
	result interface{},
	messages []message,
	status int,
) {
	bytes, err := json.Marshal(newResponse(result, messages))
	if err != nil {
		logger.Error(err)
		return
	}

	resp.WriteHeader(status)
	resp.Write(bytes)
}
