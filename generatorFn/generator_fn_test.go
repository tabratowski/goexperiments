package generatorFn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func genFn(yield func(res Result[int])) {
	for i := range 10 {
		yield(Result[int]{Value: i})
		fmt.Println(i)
	}
	fmt.Println("done")
}

func genFn2(yield func(res Result[int])) {
	yield(Result[int]{Value: 10})
	yield(Result[int]{Value: 20})
	yield(Result[int]{Value: 30})
}

func TestNewGeneratorFn(t *testing.T) {
	for _, test := range []struct {
		name        string
		genFn       func(yield func(res Result[int]))
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
