package giphy

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/wexel-nath/meat-night/pkg/config"
)

const (
	giphyBaseURL = "http://api.giphy.com/v1/gifs"
)

type data struct {
	EmbedUrl         string `json:"embed_url"`
	ImageOriginalUrl string `json:"image_original_url"`
}

type giphyResult struct {
	Data data `json:"data"`
}

// GetRandomGif returns an embeddable link of a random gif
func GetRandomGif(tag string) (string, error) {
	apiKey := config.GetGiphyApiKey()
	if apiKey == "" {
		return "", errors.New("no giphy api key is set")
	}

	url := fmt.Sprintf("%s/random?api_key=%s&tag=%s", giphyBaseURL, apiKey, tag)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{ Timeout: 10 * time.Second }

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	var result giphyResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return result.Data.ImageOriginalUrl, nil
}
