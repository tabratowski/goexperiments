package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinLength(t *testing.T) {
	data := &testStruct{
		someFloat: 2.2,
		name:      "testStruct",
		arr:       []int{1, 2, 3, 4},
		m:         map[int]string{1: "one", 2: "two", 3: "three"},
	}
	for _, test := range []struct {
		name           string
		fieldValue     hasLength
		fieldName      string
		isValid        bool
		minLength      int
		expectedErrors []error
	}{
		{
			name:           "TestMinLength - string too short",
			fieldName:      "Name",
			fieldValue:     VString(data.name),
			isValid:        false,
			minLength:      15,
			expectedErrors: []error{ErrMinLength("Name", 15)},
		},
		{
			name:       "TestMinLength - string ok",
			fieldValue: VString(data.name),
			isValid:    true,
			minLength:  3,
		},
		{
			name:           "TestMinLength - array too short",
			fieldName:      "Arr",
			fieldValue:     VArray(data.arr),
			isValid:        false,
			minLength:      5,
			expectedErrors: []error{ErrMinLength("Arr", 5)},
		},
		{
			name:       "TestMinLength - array ok",
			fieldName:  "Arr",
			fieldValue: VArray(data.arr),
			isValid:    true,
			minLength:  3,
		},
		{
			name:           "TestMinLength - map too short",
			fieldName:      "M",
			fieldValue:     VMap(data.m),
			isValid:        false,
			minLength:      5,
			expectedErrors: []error{ErrMinLength("M", 5)},
		},
		{
			name:       "TestMinLength - map ok",
			fieldName:  "M",
			fieldValue: VMap(data.m),
			isValid:    true,
			minLength:  3,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(test.fieldValue, test.fieldName).Rules(MinLength(test.minLength)))
			// Act
			errs := validator.Validate()
			// Assert
			assert.Equal(t, test.expectedErrors, errs)
		})
	}
}

func TestMaxLength(t *testing.T) {
	data := &testStruct{
		someFloat: 2.2,
		someVal:   1,
		name:      "testStruct",
		arr:       []int{1, 2, 3, 4},
		m:         map[int]string{1: "one", 2: "two", 3: "three"},
	}
	for _, test := range []struct {
		name           string
		fieldValue     hasLength
		fieldName      string
		isValid        bool
		maxLength      int
		expectedErrors []error
	}{
		{
			name:           "TestMaxLength - string too long",
			fieldName:      "Name",
			fieldValue:     VString(data.name),
			isValid:        false,
			maxLength:      5,
			expectedErrors: []error{ErrMaxLength("Name", 5)},
		},
		{
			name:       "TestMaxLength - string ok",
			fieldValue: VString(data.name),
			isValid:    true,
			maxLength:  13,
		},
		{
			name:           "TestMaxLength - array too long",
			fieldName:      "Arr",
			fieldValue:     VArray(data.arr),
			isValid:        false,
			maxLength:      2,
			expectedErrors: []error{ErrMaxLength("Arr", 2)},
		},
		{
			name:       "TestMaxLength - array ok",
			fieldName:  "Arr",
			fieldValue: VArray(data.arr),
			isValid:    true,
			maxLength:  13,
		},
		{
			name:           "TestMaxLength - map too long",
			fieldName:      "M",
			fieldValue:     VMap(data.m),
			isValid:        false,
			maxLength:      2,
			expectedErrors: []error{ErrMaxLength("M", 2)},
		},
		{
			name:       "TestMaxLength - map ok",
			fieldName:  "M",
			fieldValue: VMap(data.m),
			isValid:    true,
			maxLength:  3,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(test.fieldValue, test.fieldName).Rules(MaxLength(test.maxLength)))
			// Act
			errs := validator.Validate()
			// Assert
			assert.Equal(t, test.expectedErrors, errs)
		})
	}
}
