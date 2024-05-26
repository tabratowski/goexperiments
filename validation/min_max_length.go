package validation

import (
	"fmt"
)

type ErrMinMaxLengthFn func(fieldName string, val int) error

// There is no way (Or I can't figure it out ;p) to create generic constraint which will check if given type has length - using interface
type hasLength interface {
	Len() int
}

// lenght validation wrapper structs
type vString struct {
	string
}

func VString(str string) hasLength {
	return vString{
		string: str,
	}
}

func (v vString) Len() int {
	return len(v.string)
}

type vArray[T any] struct {
	arr []T
}

func VArray[T any](arr []T) hasLength {
	return vArray[T]{
		arr: arr,
	}
}

func (v vArray[T]) Len() int {
	return len(v.arr)
}

type vMap[T comparable, K any] struct {
	m map[T]K
}

func VMap[T comparable, K any](m map[T]K) hasLength {
	return vMap[T, K]{
		m: m,
	}
}

func (v vMap[T, K]) Len() int {
	return len(v.m)
}

type minLengthRule[T hasLength] struct {
	f         *field[hasLength]
	minLength int
	errFn     ErrMinMaxLengthFn
}

func MinLength(minLength int) RuleFn[hasLength] {
	return func(f *field[hasLength]) ValidationRule {
		return &minLengthRule[hasLength]{
			minLength: minLength,
			f:         f,
			errFn:     ErrMinLength,
		}
	}
}

func ErrMinLength(fieldName string, minLength int) error {
	return fmt.Errorf("value %s length must be greater or equal than %d", fieldName, minLength)
}

func (r *minLengthRule[T]) Validate() (bool, error) {
	if r.f.value.Len() >= r.minLength {
		return true, nil
	}
	return false, r.errFn(r.f.name, r.minLength)
}

type maxLengthRule[T hasLength] struct {
	f         *field[T]
	maxLength int
	errFn     ErrMinMaxLengthFn
}

func MaxLength(maxLength int) RuleFn[hasLength] {
	return func(f *field[hasLength]) ValidationRule {
		return &maxLengthRule[hasLength]{
			f:         f,
			maxLength: maxLength,
			errFn:     ErrMaxLength,
		}
	}
}

func (r *maxLengthRule[T]) Validate() (bool, error) {
	if r.f.value.Len() <= r.maxLength {
		return true, nil
	}
	return false, r.errFn(r.f.name, r.maxLength)
}

func ErrMaxLength(fieldName string, maxLength int) error {
	return fmt.Errorf("value %s length must be lesser or equal than %d", fieldName, maxLength)
}
