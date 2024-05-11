package calc

import (
	"os"
)

func ReadFile(filepath string) ([]byte, error) {
	body, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	return body, nil
}
