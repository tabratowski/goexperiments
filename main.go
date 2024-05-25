package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/tabratowski/gotimeout/timeout"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	ctx := context.Background()
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			err := timeout.Do(ctx, time.Duration(rand.Intn(100)+50), func(ctx context.Context) error {
				for i := 0; i < 2; i++ {
					if ctx.Err() != nil {
						fmt.Println("Cancelled func", ctx.Err())
						return fmt.Errorf("cancelled from fn")
					}
					time.Sleep(time.Duration(rand.Intn(60) + 25))
				}
				return nil
			})
			fmt.Println(err)
		}(i)
	}
	wg.Wait()
}
