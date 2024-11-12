package generatorFn

import (
	"fmt"
	"sync/atomic"
)

type Result[T any] struct {
	Err   error
	Value T
}

type SafeChannel[T any] struct {
	ch     chan Result[T]
	closed int32
}

func NewSafeChannel[T any]() *SafeChannel[T] {
	return &SafeChannel[T]{
		ch: make(chan Result[T]),
	}
}

func (c *SafeChannel[T]) Send(value T) bool {
	if atomic.LoadInt32(&c.closed) == 1 {
		return false
	}
	c.ch <- Result[T]{Value: value}
	return true
}

func (c *SafeChannel[T]) Close() {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		close(c.ch)
	}
}

func (c *SafeChannel[T]) Receive() (Result[T], bool) {
	r, ok := <-c.ch
	return r, ok
}

type GeneratorFn[T any] struct {
	resCh *SafeChannel[T]
	fn    func(yield func(result Result[T]) bool)
}

func NewGeneratorFn[T any](fn func(yield func(result Result[T]) bool)) *GeneratorFn[T] {
	return &GeneratorFn[T]{resCh: NewSafeChannel[T](), fn: fn}
}

func (gen *GeneratorFn[T]) Start() {
	go func() {
		gen.fn(func(result Result[T]) bool {
			fmt.Printf("result no: %v\r\n", result.Value)
			if !gen.resCh.Send(result.Value) {
				return false
			}
			return true
		})
		gen.Close()
	}()
}
func (gen *GeneratorFn[T]) Next() (Result[T], bool) {
	res, ok := gen.resCh.Receive()
	if ok {
		return res, true
	}
	return *new(Result[T]), false
}

func (gen *GeneratorFn[T]) Close() {
	gen.resCh.Close()
}
