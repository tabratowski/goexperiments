package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMixedRulesString(t *testing.T) {
	for _, test := range []struct {
		name           string
		fieldName      string
		testStruct     testStruct
		expectedErrors []error
		ruleSet        []RuleFn[hasLength]
	}{
		{
			name:       "Required, minLength 10 maxLength 20",
			testStruct: testStruct{name: "some name 2", surname: "surname 45"},
			ruleSet:    []RuleFn[hasLength]{Required[hasLength], MinLength(10), MaxLength(20)},
		},
		{
			name:           "Required, minLength 10 maxLength 20 - expected min length error",
			testStruct:     testStruct{name: "some name", surname: "some surname"},
			fieldName:      "NAME",
			ruleSet:        []RuleFn[hasLength]{Required[hasLength], MinLength(10), MaxLength(20)},
			expectedErrors: []error{errors.New("value NAME length must be greater or equal than 10")},
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			// Arrange && Act
			errs := RuleSet(
				Field(VString(test.testStruct.name), test.fieldName).Rules(test.ruleSet...),
				Field(VString(test.testStruct.surname), test.fieldName).Rules(test.ruleSet...),
			).Validate()

			// Assert
			assert.Equal(t, test.expectedErrors, errs)
		})
	}
}
