package client

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

func BuildRequest(r models.Request) (*http.Request, error) {

	var body io.Reader

	if r.Body != "" {
		body = strings.NewReader(r.Body)
	}

	req, err := http.NewRequest(
		r.Method,
		r.URL,
		body,
	)

	if err != nil {
		return nil, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	// Default content type for body
	if r.Body != "" &&
		req.Header.Get("Content-Type") == "" {

		req.Header.Set(
			"Content-Type",
			"application/json",
		)
	}

	return req, nil
}

func Do(
	client *http.Client,
	request models.Request,

) (*http.Response, time.Duration, error) {

	req, err := BuildRequest(request)

	if err != nil {
		return nil, 0, err
	}

	start := time.Now()
	resp, err := client.Do(req)

	duration := time.Since(start)

	if err != nil {
		return nil, duration, err
	}

	return resp, duration, nil
}
