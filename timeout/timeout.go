package timeout

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

func Do(ctx context.Context, d time.Duration, fn func(ctx context.Context) error) error {
	done := make(chan struct{})
	defer close(done)
	var err error
	mu := &sync.Mutex{}
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	go func() {
		mu.Lock()
		err = fn(ctx)
		defer mu.Unlock()
		if ctx.Err() != nil {
			return
		}
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return errors.New("context has been cancelled")
	case <-done:
		fmt.Println("done")
		return err
	}
}
