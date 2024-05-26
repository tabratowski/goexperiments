package validation

import (
	"fmt"
)

type ErrRequiredFn func(fieldName string) error

type requiredRule[T any] struct {
	f     *field[T]
	errFn ErrRequiredFn
}

func Required[T any](f *field[T]) ValidationRule {
	return &requiredRule[T]{
		f:     f,
		errFn: ErrRequired,
	}
}

func (v *requiredRule[T]) Validate() (bool, error) {
	if empty := isEmpty(v.f.value); !empty {
		return true, nil
	}
	return false, v.errFn(v.f.name)
}

func ErrRequired(fieldName string) error {
	return fmt.Errorf("value %s, cannot be nil or empty", fieldName)
}
