package validation

type RuleFn[T any] func(*field[T]) ValidationRule

type ruleField interface {
	RulesSet() []ValidationRule
	Errors() []error
}

type field[T any] struct {
	value  T
	name   string
	rules  []ValidationRule
	errors []error
}

func Field[T any](f T, name string) *field[T] {
	return &field[T]{
		value: f,
		name:  name,
	}
}

func (f *field[T]) Rules(ruleFns ...RuleFn[T]) *field[T] {
	for _, ruleFn := range ruleFns {
		f.rules = append(f.rules, ruleFn(f))
	}
	return f
}

func (f *field[T]) RulesSet() []ValidationRule {
	return f.rules
}

func (f *field[T]) Errors() []error {
	return f.errors
}
