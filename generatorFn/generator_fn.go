package generatorFn

import (
	"fmt"
)

type Result[T any] struct {
	Err   error
	Value T
}

type GeneratorFn[T any] struct {
	resCh chan Result[T]
	fn    func(yield func(result Result[T]))
}

func NewGeneratorFn[T any](fn func(yield func(result Result[T]))) *GeneratorFn[T] {
	return &GeneratorFn[T]{resCh: make(chan Result[T]), fn: fn}
}

func (gen *GeneratorFn[T]) Start() {
	go func() {
		gen.fn(func(result Result[T]) {
			fmt.Printf("result no: %v\r\n", result.Value)
			gen.resCh <- result
		})
		close(gen.resCh)
	}()
}
func (gen *GeneratorFn[T]) Next() (Result[T], bool) {
	res, ok := <-gen.resCh
	if ok {
		return res, true
	}
	return *new(Result[T]), false
}
