package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	data := &testStruct{
		name:      "testStruct",
		surname:   "testsurname",
		someVal:   8,
		someFloat: 12.5,
		sub: &testSub{
			subName: "subname",
		},
	}

	for _, test := range []struct {
		name          string
		pattern       string
		isMatch       bool
		expectedError error
	}{
		{
			name:    "Match - matches pattern",
			isMatch: true,
			pattern: `.+estStr.+`,
		},
		{
			name:          "Match -does not match pattern",
			isMatch:       false,
			pattern:       `.+esttttStr.+`,
			expectedError: ErrMatch("testStruct", "Name", `.+esttttStr.+`),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(data.name, "Name").Rules(Match(test.pattern)))

			// Act
			errs := validator.Validate()

			// Assert
			if len(errs) > 0 {
				assert.Equal(t, test.expectedError, errs[0])
			}
		})
	}
}
