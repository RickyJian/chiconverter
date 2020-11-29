package utils

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ReadAll return all text from reader
func ReadAll(r io.Reader) (string, error) {
	var builder strings.Builder
	reader := bufio.NewReader(r)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return "", fmt.Errorf("falied to read text: %v", err)
		}
		builder.WriteString(string(line))
		builder.WriteString("\n")
	}
	return builder.String(), nil
}
