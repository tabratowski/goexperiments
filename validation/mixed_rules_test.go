package validation

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMixedRulesString(t *testing.T) {
	for _, test := range []struct {
		name           string
		fieldName      string
		fieldValue     string
		expectedErrors []error
		ruleSet        []RuleFn[hasLength]
	}{
		{
			name:       "Required, minLength 10 maxLength 20",
			fieldValue: "some name 2",
			ruleSet:    []RuleFn[hasLength]{Required[hasLength], MinLength(10), MaxLength(20)},
		},
		{
			name:           "Required, minLength 10 maxLength 20 - expected min length error",
			fieldValue:     "some name",
			fieldName:      "NAME",
			ruleSet:        []RuleFn[hasLength]{Required[hasLength], MinLength(10), MaxLength(20)},
			expectedErrors: []error{errors.New("value NAME length must be greater or equal than 10")},
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			// Arrange && Act
			r := reflect.TypeOf(test.fieldValue)
			fmt.Print(r.Implements(reflect.TypeOf((*fmt.Stringer)(nil)).Elem()))
			errs := RuleSet(Field(VString(test.fieldValue), test.fieldName).Rules(test.ruleSet...)).Validate()

			// Assert
			assert.Equal(t, test.expectedErrors, errs)
		})
	}
}
