package formatter

import (
	"fmt"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

type TableFormatter struct{}

func (TableFormatter) Print(response models.Response) error {

	fmt.Println("Status :", response.Status)

	fmt.Println("Code   :", response.StatusCode)

	fmt.Println("Time   :", response.Duration)

	fmt.Println()

	fmt.Println("Headers")

	for key, values := range response.Headers {

		for _, value := range values {

			fmt.Printf("%s : %s\n", key, value)

		}

	}

	fmt.Println()

	fmt.Println(string(response.Body))

	return nil

}
