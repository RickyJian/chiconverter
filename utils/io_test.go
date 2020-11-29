package utils

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAll(t *testing.T) {
	var tests = []*struct {
		text     string
		expected string
	}{
		{
			text: `天地玄黃，宇宙洪荒。`,
			expected: fmt.Sprintf("%s\n", "天地玄黃，宇宙洪荒。"),
		},
		{
			text: `天地玄黃，宇宙洪荒。
日月盈昃，辰宿列張。
寒來暑往，秋收冬藏。`,
			expected: fmt.Sprintf("%s\n%s\n%s\n", "天地玄黃，宇宙洪荒。", "日月盈昃，辰宿列張。", "寒來暑往，秋收冬藏。"),
		},
	}
	for _, test := range tests {
		reader := bufio.NewReader(strings.NewReader(test.text))
		text, _ := ReadAll(reader)
		assert.Equal(t, test.expected, text)
	}
}
