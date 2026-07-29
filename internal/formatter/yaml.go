package formatter

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
	"go.yaml.in/yaml/v3"
)

type YAMLFormatter struct{}

func (YAMLFormatter) Print(response models.Response) error {

	data, err := yaml.Marshal(response)

	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
