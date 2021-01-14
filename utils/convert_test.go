package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockConvert struct{}

func (m *mockConvert) Convert() (map[int]string, error) {
	return map[int]string{0: "mock convert"}, nil
}

func TestConvert(t *testing.T) {
	var tests = []*struct {
		converter   Converter
		expected    map[int]string
		expectedErr error
	}{
		{
			converter:   new(mockConvert),
			expected:    map[int]string{0: "mock convert"},
			expectedErr: nil,
		},
	}
	for _, test := range tests {
		result, err := Convert(test.converter)
		assert.Equal(t, test.expected, result)
		assert.Equal(t, test.expectedErr, err)
	}
}


func TestBatch(t *testing.T) {
	mockConvertFunc := func(in string) (string, error) {
		return in, nil
	}
	var tests = []*struct {
		src  string
	}{
		{
			src: "../test/thousand_character_classic_cht",
		},
		{
			src: "../test/three_character_classic_cht",
		},
	}
	for _, test := range tests {
		bs, _ := ReadAll(test.src)
		results, _ := Batch(bs, mockConvertFunc)
		for i := 0; i < len(results); i++ {
			assert.Equal(t, string(bs[i]), results[i])
		}
	}
}
