package timeout

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			err := Do(time.Duration(rand.Intn(100)+50)*time.Millisecond, func() Result[int] {
				for i := 0; i < 1; i++ {
					time.Sleep(time.Duration(rand.Intn(100)+25) * time.Millisecond)
				}
				return Result[int]{Value: 4}
			})
			fmt.Println(err)
		}(i)
	}
	wg.Wait()
}
