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
	ResponseBody  []byte
}

func ApiAnalyzer(method string, URL string, body string) (Api, error) {

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

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(normalizedMethod, normalizedUrl, strings.NewReader(body))
	if err != nil {
		return Api{}, err
	}

	if normalizedMethod == http.MethodPost || normalizedMethod == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return Api{}, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return Api{}, err
	}

	api.StatusCode = resp.StatusCode
	api.StatusMessage = resp.Status
	api.ResponseTime = time.Since(start)
	api.ContentType = resp.Header.Get("Content-Type")
	api.BodySize = float64(len(responseBody))
	api.ResponseBody = responseBody

	return api, nil
}
