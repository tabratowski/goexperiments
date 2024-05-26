package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	data := &testStruct{
		someVal: 8,
	}

	for _, test := range []struct {
		name          string
		min           int
		ok            bool
		expectedError error
	}{
		{
			name:          "TestMin - int too low",
			min:           10,
			ok:            false,
			expectedError: ErrMin("SomeVal", 10),
		},
		{
			name: "TestMin - int ok",
			min:  5,
			ok:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(data.someVal, "SomeVal").Rules(Min(test.min)))

			// Act
			errs := validator.Validate()

			// Assert
			if len(errs) > 0 {
				assert.Equal(t, test.expectedError, errs[0])
			}
		})
	}
}

func TestMaxInt(t *testing.T) {
	data := &testStruct{
		someVal: 8,
	}

	for _, test := range []struct {
		name          string
		max           int
		ok            bool
		expectedError error
	}{
		{
			name:          "TestMax - int too high",
			max:           5,
			ok:            false,
			expectedError: ErrMax("SomeVal", 5),
		},
		{
			name: "TestMax - int ok",
			max:  10,
			ok:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(data.someVal, "SomeVal").Rules(Max(test.max)))

			// Act
			errs := validator.Validate()

			// Assert
			if len(errs) > 0 {
				assert.Equal(t, test.expectedError, errs[0])
			}
		})
	}
}

func TestMinFloat(t *testing.T) {
	data := &testStruct{
		someFloat: 12.5,
	}

	for _, test := range []struct {
		name          string
		min           float64
		ok            bool
		expectedError error
	}{
		{
			name:          "TestMin - float too low",
			min:           12.7,
			ok:            false,
			expectedError: ErrMin("SomeFloat", 12.7),
		},
		{
			name: "TestMin - float ok",
			min:  5.3,
			ok:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(data.someFloat, "SomeFloat").Rules(Min(test.min)))

			// Act
			errs := validator.Validate()

			// Assert
			if len(errs) > 0 {
				assert.Equal(t, test.expectedError, errs[0])
			}
		})
	}
}

func TestMaxFloat(t *testing.T) {
	data := &testStruct{
		someFloat: 12.5,
	}

	for _, test := range []struct {
		name          string
		max           float64
		ok            bool
		expectedError error
	}{
		{
			name:          "TestMax - float too high",
			max:           12.3,
			ok:            false,
			expectedError: ErrMax("SomeFloat", 12.3),
		},
		{
			name: "TestMax - float ok",
			max:  15.3,
			ok:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(data.someFloat, "SomeFloat").Rules(Max(test.max)))

			// Act
			errs := validator.Validate()

			// Assert
			if len(errs) > 0 {
				assert.Equal(t, test.expectedError, errs[0])
			}
		})
	}
}

func TestMinMaxFloat(t *testing.T) {
	data := &testStruct{
		someFloat: 12.8,
	}

	for _, test := range []struct {
		name          string
		min           float64
		max           float64
		ok            bool
		expectedError error
	}{
		{
			name:          "TestMinMax - float too low",
			min:           12.9,
			max:           13.0,
			ok:            false,
			expectedError: ErrMin("SomeFloat", 12.9),
		},
		{
			name:          "TestMinMax - float too high",
			min:           12.3,
			max:           12.4,
			ok:            false,
			expectedError: ErrMax("SomeFloat", 12.4),
		},
		{
			name: "TestMinMax - float ok",
			min:  5.3,
			max:  15.3,
			ok:   true,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			validator := RuleSet(Field(data.someFloat, "SomeFloat").Rules(Min(test.min), Max(test.max)))

			// Act
			errs := validator.Validate()

			// Assert
			if len(errs) > 0 {
				assert.Equal(t, test.expectedError, errs[0])
			}
		})
	}
}
