package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockConvert struct{}

func (m *mockConvert) Convert() (string, error) {
	return "mock convert", nil
}

type mockConvertUnimplemented struct{}

func TestConvert(t *testing.T) {
	var tests = []*struct {
		converter   interface{}
		expected    string
		expectedErr error
	}{
		{
			converter:   new(mockConvert),
			expected:    "mock convert",
			expectedErr: nil,
		},
		{
			converter:   new(mockConvertUnimplemented),
			expected:    "",
			expectedErr: ErrConvertUnimplemented,
		},
	}
	for _, test := range tests {
		result, err := Convert(test.converter)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expectedErr, err)
	}
}
