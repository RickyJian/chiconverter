package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	underscore = "_"
	timeLayout = "2006-01-02_15:04:05"
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
	// ErrSourceIsNotFile describes source is not file
	ErrSourceIsNotFile = errors.New("source is not file")
	// ErrDestinationIsNotDir describes is not directory
	ErrDestinationIsNotDir = errors.New("destination is not directory")
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
	} else if mode := fileStat.Mode(); !mode.IsRegular() {
		return nil, ErrSourceIsNotFile
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

// DestFileName returns destination file name
func DestFileName(src, dest string) (string, error) {
	if fileStat, err := os.Stat(dest); err != nil {
		return "", fmt.Errorf("invalid path: %w", err)
	} else if mode := fileStat.Mode(); mode.IsRegular() || !fileStat.IsDir() {
		return "", ErrDestinationIsNotDir
	}

	fileExt := filepath.Ext(src)
	var builder strings.Builder
	builder.WriteString(filepath.Base(src[:len(src)-len(fileExt)]))
	builder.WriteString(underscore)
	builder.WriteString(time.Now().Format(timeLayout))
	if fileExt != "" {
		builder.WriteString(fileExt)
	}
	path, _ := filepath.Abs(filepath.Join(dest, builder.String()))
	return path, nil
}
