package common_test

import (
	"testing"

	"github.com/elC0mpa/gonet/common"
	"github.com/stretchr/testify/assert"
)

func TestFromBytesToString(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{name: "Formatting when input is bytes", input: 1023, expected: "1023 B"},
		{name: "Formatting when input is kilo bytes", input: 2048, expected: "2 KB"},
		{name: "Formatting when input is mega bytes", input: 1572864, expected: "1.5 MB"},
		{name: "Formatting when input is gigabytes", input: 1610612736, expected: "1.5 GB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := common.FromBytesToString(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
