package tdd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type glober struct {
	name     string
	position string
}

func NewGlober() glober {

	glober := glober{"timo", "developer"}

	return glober
}

func TestNewGlober(t *testing.T) {
	glober := NewGlober()
	assert.EqualValues(t, "timo", glober.name)

}

func Invert(input bool) bool {
	return !input
}

func TestInvert(t *testing.T) {
	testCases := []struct {
		name           string
		input          bool
		expectedResult bool
	}{
		{
			name:           "validate is true on false input",
			input:          false,
			expectedResult: true,
		},
		{
			name:           "validate is false on true input",
			input:          true,
			expectedResult: false,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare
			//Nothing to prepare in this case

			// Act
			result := Invert(tt.input)

			// Assert
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
