package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzNumber(t *testing.T) {

	testCases := []struct {
		name           string
		input          int
		expectedResult string
	}{
		{
			name:           "Case of Fizz alone",
			input:          3,
			expectedResult: fizzRes,
		},

		{
			name:           "Case of Buzz alone",
			input:          5,
			expectedResult: buzzRes,
		},

		{
			name:           "Case of FizzBuzz",
			input:          15,
			expectedResult: FizzBuzzRes,
		},

		{
			name:           "Case of not Fizz nor Buzz",
			input:          4,
			expectedResult: "4",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := FizzBuzz(tc.input)

			assert.Equal(t, tc.expectedResult, result)

		})
	}

}
