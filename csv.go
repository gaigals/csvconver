package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseCSV(input string) (map[string][]string, error) {
	if isInputFile(input) {
		file, err := readFile(input)
		if err != nil {
			return nil, err
		}

		defer file.Close()
		return processCSV(bufio.NewReader(file))
	}

	return processCSV(strings.NewReader(input))
}

func isInputFile(input string) bool {
	return strings.HasSuffix(strings.ToLower(input), ".csv")
}

func readFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("file open error: %w", err)
	}

	return file, nil
}

func processCSV(reader io.Reader) (map[string][]string, error) {
	csvReader := csv.NewReader(reader)
	csvReader.LazyQuotes = true

	headers, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("csv header read error: %w", err)
	}

	data := make(map[string][]string)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("csv read error: %w", err)
		}

		for idx, colVal := range row {
			data[headers[idx]] = append(data[headers[idx]], colVal)
		}

	}

	return data, nil
}
