package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {
	var tests = []*struct {
		src         string
		expected    [][]byte
		expectedErr error
	}{
		{
			src:         "../test/file_empty",
			expected:    [][]byte{},
			expectedErr: ErrEmptyFile,
		},
		{
			src: "../test/file_line",
			expected: [][]byte{
				[]byte("天地玄黃，宇宙洪荒。"),
			},
		},
		{
			src: "../test/file_lines",
			expected: [][]byte{
				[]byte("天地玄黃，宇宙洪荒。"),
				[]byte("日月盈昃，辰宿列張。"),
				[]byte("寒來暑往，秋收冬藏。"),
			},
		},
	}
	for _, test := range tests {
		texts, err := ReadAll(test.src)
		assert.Equal(t, test.expectedErr, err)
		if err == nil {
			assert.Equal(t, test.expected, texts)
		}
	}
}
