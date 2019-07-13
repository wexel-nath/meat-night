package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/wexel-nath/meat-night/pkg/giphy"
	"github.com/wexel-nath/meat-night/pkg/logger"
)

type jsonHandlerFunc func(r *http.Request, ps httprouter.Params) (interface{}, int, error)

func jsonResponseHandler(handler jsonHandlerFunc) httprouter.Handle {
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

func writeJsonResponse(resp http.ResponseWriter, result interface{}, messages []string, status int) {
	response := struct {
		Result   interface{} `json:"result"`
		Messages []string    `json:"messages"`
	}{
		Result:   result,
		Messages: messages,
	}
	bytes, err := json.Marshal(response)
	if err != nil {
		logger.Error(err)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(status)
	resp.Write(bytes)
}

type giphyHandlerFunc func(ps httprouter.Params) (string, error)

func giphyResponseHandler(handler giphyHandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "text/html")
		tag, err := handler(ps)
		if err != nil {
			logger.Error(err)
			//tag = "something-went-wrong"
		}

		w.Write([]byte(buildGiphyHtml(tag)))
	}
}

func buildGiphyHtml(tag string) string {
	giphyUrl, err := giphy.GetRandomGif(tag)
	if err != nil {
		logger.Error(err)
		return ""
	}

	format := `
		<html>
			<img src="%s" />
		</html>
	`

	return fmt.Sprintf(format, giphyUrl)
}
