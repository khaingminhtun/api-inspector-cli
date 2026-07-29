package formatter

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

type Formatter interface {
	Print(models.Response) error
}

func Print(output string, response models.Response) error {

	switch output {

	case "json":
		return JSONFormatter{}.Print(response)

	case "yaml":
		return YAMLFormatter{}.Print(response)

	case "table":
		return TableFormatter{}.Print(response)

	default:
		return fmt.Errorf("unsupported output format: %s", output)

	}
}
