package main

import (
	"encoding/json"
	"fmt"
)

func CastToJSON(data map[string][]string, asPretty bool) (string, error) {
	if asPretty {
		bytes, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return "", fmt.Errorf("json marshal error: %w", err)
		}
		return string(bytes), nil
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("json marshal error: %w", err)
	}

	return string(bytes), nil
}
