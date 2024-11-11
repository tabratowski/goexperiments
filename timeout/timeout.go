package timeout

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Result[T any] struct {
	Value T
	Err   error
}

func Do[T any](d time.Duration, fn func() Result[T]) Result[T] {
	resultCh := make(chan Result[T], 1)
	var chClosed bool
	mu := sync.Mutex{}
	defer func() {
		mu.Lock()
		defer mu.Unlock()
		chClosed = true
		close(resultCh)
	}()
	go func() {
		mu.Lock()
		defer mu.Unlock()
		res := fn()
		if !chClosed {
			resultCh <- res
		}
	}()
	select {
	case <-time.After(d):
		return Result[T]{Err: errors.New("timeout occurred")}
	case res := <-resultCh:
		fmt.Println("resultCh")
		return res
	}
}
