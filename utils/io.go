package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	// MaxFileSize defines max size for reading file
	MaxFileSize int64 = 1024 * 1024 * 1024
)

var (
	// ErrFileOutOfSize describes file out of size
	ErrFileOutOfSize = errors.New("file out of size")
	// ErrEmptyFile describes empty file
	ErrEmptyFile = errors.New("empty file")
)

// ReadAll read all text from file
func ReadAll(src string) ([][]byte, error) {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	} else if size := fileStat.Size(); size > MaxFileSize {
		return nil, ErrFileOutOfSize
	} else if size == 0 {
		return nil, ErrEmptyFile
	}

	reader := bufio.NewReader(file)
	var text [][]byte
	for {
		words, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("falied to read text: %w", err)
		}
		t := make([]byte, len(words))
		// use `copy` to prevent first index of `text` change
		// to last index of bytes when read the last line
		copy(t, words)
		text = append(text, t)
	}
	return text, nil
}
