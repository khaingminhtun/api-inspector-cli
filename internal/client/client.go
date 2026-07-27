package client

import (
	"net/http"
	"time"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

func New(timeout int) *http.Client{
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}

func MakeRequest(
	request models.Request,
	timeout int,
) (models.Response, error){

	httpClient := New(timeout)

	resp, duration, err := Do(
		httpClient,
		request,
	)

	if err != nil {
       return models.Response{},err
	}

	return ParseResponse(
		resp,
		duration,
	)
}