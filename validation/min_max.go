package validation

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type ErrMinMaxFn func(string, any) error

type numeric interface {
	constraints.Float | constraints.Integer
}

type minRule[T numeric] struct {
	minValue T
	f        *field[T]
	errFn    ErrMinMaxFn
}

func Min[T numeric](min T) RuleFn[T] {
	return func(f *field[T]) ValidationRule {
		return &minRule[T]{
			f:        f,
			minValue: min,
			errFn:    ErrMin,
		}
	}
}

func (v *minRule[T]) Validate() (bool, error) {
	if v.f.value >= v.minValue {
		return true, nil
	}
	return false, fmt.Errorf("value %s value must be greater or equal than %v", v.f.name, v.minValue)
}

func ErrMin(fieldName string, minValue any) error {
	return fmt.Errorf("value %s value must be greater or equal than %v", fieldName, minValue)
}

type maxRule[T numeric] struct {
	f        *field[T]
	maxValue T
	errFn    ErrMinMaxFn
}

func Max[T numeric](max T) RuleFn[T] {
	return func(f *field[T]) ValidationRule {
		return &maxRule[T]{
			f:        f,
			maxValue: max,
			errFn:    ErrMax,
		}
	}
}

func (v *maxRule[T]) Validate() (bool, error) {
	if v.f.value <= v.maxValue {
		return true, nil
	}
	return false, fmt.Errorf("value %s value must be lesser or equal than %v", v.f.name, v.maxValue)
}

func ErrMax(fieldName string, maxValue any) error {
	return fmt.Errorf("value %s value must be lesser or equal than %v", fieldName, maxValue)
}
