package client

import (
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

	return models.Response{

		StatusCode: resp.StatusCode,

		Status: resp.Status,

		Headers: resp.Header,

		Body: body,

		Duration: duration,
	}, nil
}
