package module

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Api struct {
	StatusCode    int
	StatusMessage string
	ResponseTime  time.Duration
	ContentType   string
	BodySize      float64
}

func ApiAnalyzer(method string, URL string) (Api, error) {

	var api Api

	normalizedMethod := strings.ToUpper(strings.TrimSpace(method))

	switch normalizedMethod {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		return Api{}, errors.New("invalid api method")
	}

	normalizedUrl := strings.TrimSpace(URL)
	if normalizedUrl == "" {
		return Api{}, errors.New("invalid api url")
	}

	u, err := url.Parse(normalizedUrl)
	if err != nil {
		return Api{}, err
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return Api{}, errors.New("invalid api url scheme")
	}

	if u.Host == "" {
		return Api{}, errors.New("invalid api url host")
	}

	client := &http.Client{}

	req, err := http.NewRequest(normalizedMethod, normalizedUrl, nil)
	if err != nil {
		return Api{}, err
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return Api{}, err
	}
	defer resp.Body.Close()

	size, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return Api{}, err
	}

	api.StatusCode = resp.StatusCode
	api.StatusMessage = resp.Status
	api.ResponseTime = time.Since(start)
	api.ContentType = resp.Header.Get("Content-Type")
	api.BodySize = float64(size)

	return api, nil
}
