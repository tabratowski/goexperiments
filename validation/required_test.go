package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredString(t *testing.T) {
	for _, test := range []struct {
		name           string
		fieldValue     string
		fieldName      string
		isValid        bool
		expectedErrors []error
	}{
		{
			name:           "TestRequired - name is empty",
			fieldValue:     "",
			isValid:        false,
			expectedErrors: []error{ErrRequired("Name")},
			fieldName:      "Name",
		},
		{
			name:       "TestRequired - name is not empty",
			fieldValue: "name",
			isValid:    true,
			fieldName:  "Name",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(test.fieldValue, test.fieldName).Rules(Required[string]))

			// Act
			errs := validator.Validate()

			// Assert
			assert.Equal(t, test.expectedErrors, errs)
		})
	}
}
