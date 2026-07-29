package client

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

func ParseResponse(
	resp *http.Response,
	duration time.Duration,
) (models.Response, error) {

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return models.Response{}, err
	}

	var bodyData interface{}

	err = json.Unmarshal(body, &bodyData)

	if err != nil {
		bodyData = string(body)
	}

	return models.Response{

		StatusCode: resp.StatusCode,

		Status: resp.Status,

		Headers: resp.Header,

		Body: bodyData,

		Duration: duration.String(),
	}, nil
}
