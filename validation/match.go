package validation

import (
	"fmt"
	"regexp"
)

type strings interface {
	~string | ~byte
}

type ErrMatchFn[T strings] func(val T, fieldName string, pattern string) error

type matchRule[T strings] struct {
	pattern T
	field   *field[T]
	errFn   ErrMatchFn[T]
}

func Match[T strings](pattern T) RuleFn[T] {
	return func(f *field[T]) ValidationRule {
		return &matchRule[T]{
			field:   f,
			pattern: pattern,
			errFn:   ErrMatch[T],
		}
	}
}

func (v *matchRule[T]) Validate() (bool, error) {
	if ok, err := regexp.MatchString(string(v.pattern), string(v.field.value)); ok && err == nil {
		return true, nil
	}
	return false, v.errFn(v.field.value, v.field.name, string(v.pattern))
}

func ErrMatch[T strings](val T, fieldName string, pattern string) error {
	return fmt.Errorf("value %s value %s must match pattern %s", fieldName, string(val), pattern)
}
