package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

type JSONFormatter struct{}

func (JSONFormatter) Print(response models.Response) error {

	data, err := json.MarshalIndent(
		response,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil

}