package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/tabratowski/goexperiments/timeout"
)

func main() {
	res := timeout.Do(1*time.Second, func() timeout.Result[int] {
		time.Sleep(800 * time.Millisecond)
		return timeout.Result[int]{Value: 1}
	})
	fmt.Println(res)
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			err := timeout.Do(time.Duration(rand.Intn(100)+50)*time.Millisecond, func() timeout.Result[int] {
				for i := 0; i < 1; i++ {
					time.Sleep(time.Duration(rand.Intn(100)+25) * time.Millisecond)
				}
				return timeout.Result[int]{Value: 4}
			})
			fmt.Println(err)
		}(i)
	}
	wg.Wait()
}
