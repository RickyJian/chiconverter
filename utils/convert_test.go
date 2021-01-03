package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockConvert struct{}

func (m *mockConvert) Convert() (string, error) {
	return "mock convert", nil
}

func TestConvert(t *testing.T) {
	var tests = []*struct {
		converter   Converter
		expected    string
		expectedErr error
	}{
		{
			converter:   new(mockConvert),
			expected:    "mock convert",
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		result, err := Convert(test.converter)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expectedErr, err)
	}
}
