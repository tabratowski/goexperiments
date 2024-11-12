package generatorFn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func genFn(yield func(res Result[int]) bool) {
	for i := range 10 {
		if !yield(Result[int]{Value: i}) {
			return
		}
	}
	fmt.Println("done")
}

func genFn2(yield func(res Result[int]) bool) {
	yield(Result[int]{Value: 10})
	yield(Result[int]{Value: 20})
	yield(Result[int]{Value: 30})
}

func TestNewGeneratorFn(t *testing.T) {
	for _, test := range []struct {
		name        string
		genFn       func(yield func(res Result[int]) bool)
		expectedLen int
	}{
		{
			name:        "Gen 1",
			genFn:       genFn,
			expectedLen: 10,
		},
		{
			name:        "Gen 2",
			genFn:       genFn2,
			expectedLen: 3,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			gen := NewGeneratorFn(test.genFn)
			var allResults []int
			// Act
			gen.Start()
			defer gen.Close()
			for {
				if res, ok := gen.Next(); ok {
					allResults = append(allResults, res.Value)
					fmt.Println(res.Value)
					continue
				}
				break
			}

			// Assert
			assert.Len(t, allResults, test.expectedLen)
		})
	}
}

func TestGeneratorFn_Next(t *testing.T) {
	for _, test := range []struct {
		name           string
		nextCalls      int
		genFn          func(yield func(res Result[int]) bool)
		expectedResult Result[int]
	}{
		{
			name:           "Next - should stop on second yield",
			nextCalls:      2,
			genFn:          genFn,
			expectedResult: Result[int]{Value: 1},
		},
		{
			name:           "Next - should stop on seventh yield",
			nextCalls:      7,
			genFn:          genFn,
			expectedResult: Result[int]{Value: 6},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			gen := NewGeneratorFn(test.genFn)
			gen.Start()
			defer gen.Close()
			var result Result[int]
			for i := 0; i < test.nextCalls; i++ {
				result, _ = gen.Next()
			}
			assert.Equal(t, test.expectedResult, result)
		})
	}
}
