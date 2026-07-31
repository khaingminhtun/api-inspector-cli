package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/khaingminhtun/api-inspector-cli/internal/models"
)

func directory() string {

	home, _ := os.UserHomeDir()

	return filepath.Join(
		home,
		".apispy",
		"requests",
	)

}

func Save(
	name string,
	request models.Request,
) error {

	err := os.MkdirAll(
		directory(),
		0755,
	)

	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(
		request,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	path := filepath.Join(
		directory(),
		name+".json",
	)

	return os.WriteFile(
		path,
		data,
		0644,
	)
}

func Load(
	name string,
) (models.Request, error) {

	path := filepath.Join(
		directory(),
		name+".json",
	)

	data, err := os.ReadFile(path)

	if err != nil {
		return models.Request{}, err
	}

	var request models.Request

	err = json.Unmarshal(
		data,
		&request,
	)

	if err != nil {
		return models.Request{}, err
	}

	return request, nil
}

func List() ([]string, error) {

	files, err := os.ReadDir(directory())

	if err != nil {
		return nil, err
	}

	var requests []string

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		name := file.Name()

		if filepath.Ext(name) != ".json" {
			continue
		}

		requests = append(
			requests,
			name[:len(name)-5],
		)
	}

	return requests, nil
}

func Delete(name string) error {

	path := filepath.Join(
		directory(),
		name+".json",
	)

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return fmt.Errorf(
			"request '%s' not found",
			name,
		)
	}

	return os.Remove(path)
}
