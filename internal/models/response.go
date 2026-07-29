package models

import (
	"net/http"
)

type Response struct {
	StatusCode int `json:"status_code"`

	Status string `json:"status"`

	Headers http.Header `json:"headers"`

	Body interface{} `json:"body"`

	Duration string `json:"duration"`
}
