package main

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"math/rand"
	"sync"
	"time"
)

func main() {
	out := fetch(context.Background(), []int{1, 2, 3, 4, 5, 6}, func(ctx context.Context, in int) (*Result, error) {
		return SomeLRO(in), nil
	})
	for _, o := range out {
		fmt.Println(o.no)
	}
	fmt.Println("Len:")
	fmt.Println(len(out))
	//res := timeout.Do(1*time.Second, func() timeout.Result[int] {
	//	time.Sleep(800 * time.Millisecond)
	//	return timeout.Result[int]{Value: 1}
	//})
	//fmt.Println(res)
	//wg := &sync.WaitGroup{}
	//wg.Add(100)
	//for i := 0; i < 100; i++ {
	//	go func(i int) {
	//		defer wg.Done()
	//		err := timeout.Do(time.Duration(rand.Intn(100)+50)*time.Millisecond, func() timeout.Result[int] {
	//			for i := 0; i < 1; i++ {
	//				time.Sleep(time.Duration(rand.Intn(100)+25) * time.Millisecond)
	//			}
	//			return timeout.Result[int]{Value: 4}
	//		})
	//		fmt.Println(err)
	//	}(i)
	//}
	//wg.Wait()
	//for i, j := range TestIterator(50) {
	//	fmt.Println(i, j)
	//}
	//var slice []int
	//for i := range 50 {
	//	slice = append(slice, i)
	//}
	//for idx, i := range Filter(slice, func(i int) bool {
	//	return i%2 == 0
	//}) {
	//	fmt.Println(i, idx)
	//}
}

func TestIterator(max int) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		n := 0
		for {
			if (n%2 == 0 && !yield(n, n+1)) || n >= max {
				return
			}
			n++
		}
	}
}

func Filter[T any](items []T, f func(T) bool) iter.Seq2[T, int] {
	return func(yield func(item T, idx int) bool) {
		for idx, i := range items {
			if f(i) && !yield(i, idx+1) {
				return
			}
		}
	}
}

type Result struct{ no int }

func SomeLRO(no int) *Result {
	time.Sleep(time.Duration(rand.Intn(1000)+500) * time.Millisecond)
	return &Result{no}
}

func getResults() []*Result {
	opsCount := 10
	resultCh := make(chan *Result, opsCount)
	defer close(resultCh)
	wg := &sync.WaitGroup{}
	var r []*Result
	for i := range opsCount {
		fmt.Println(i)
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			resultCh <- SomeLRO(no)
			println(no)
		}(i)
	}
	wg.Wait()
	for range opsCount {
		r = append(r, <-resultCh)
	}
	fmt.Println(r, len(r))
	return r
}

func fetch[OUT, IN any](ctx context.Context, params []IN, fn func(ctx context.Context, in IN) (OUT, error)) []OUT {
	resultsCh := make(chan fnResult[OUT], len(params))
	defer close(resultsCh)
	var out []OUT
	for _, item := range params {
		go func(ctx context.Context, item IN) {
			r, err := fn(ctx, item)
			resultsCh <- fnResult[OUT]{r, err}
		}(ctx, item)
	}
	for range params {
		res := <-resultsCh
		if res.err != nil {
			res.err = errors.Join(res.err)
			continue
		}
		out = append(out, res.inner)
	}
	return out
}

type fnResult[T any] struct {
	inner T
	err   error
}
