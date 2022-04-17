package streams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapConcat(t *testing.T) {

	t.Run("singular flow", func(t *testing.T) {
		intMapConcat := NewMapConcat(func(i int) []int {
			out := []int{}
			for j := 0; j < i; j++ {
				out = append(out, j)
			}
			return out
		})
		expected := []int{0, 1, 2, 3, 4}
		result := []int{}

		go func() {
			intMapConcat.In <- 5
			close(intMapConcat.In)
		}()

		go intMapConcat.Run()

		for i := range intMapConcat.Out {
			result = append(result, i)
		}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("two connected flows", func(t *testing.T) {
		intMapConcat := NewMapConcat(func(i int) []int {
			out := []int{}
			for j := 0; j < i; j++ {
				out = append(out, j)
			}
			return out
		})
		intToStringMapConcat := NewMapConcat(func(i int) []string {
			return []string{fmt.Sprint(i)}
		})

		intMapConcat.Connect(intToStringMapConcat)

		expected := []string{"0", "1", "2", "3", "4"}
		result := []string{}

		go func() {
			intMapConcat.In <- 5
			close(intMapConcat.In)
		}()

		go intMapConcat.Run()
		go intToStringMapConcat.Run()

		for i := range intToStringMapConcat.Out {
			result = append(result, i)
		}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("three connected flows", func(t *testing.T) {
		intMapConcat := NewMapConcat(func(i int) []int {
			out := []int{}
			for j := 0; j < i; j++ {
				out = append(out, j)
			}
			return out
		})
		intToStringMapConcat := NewMapConcat(func(i int) []string {
			return []string{fmt.Sprint(i)}
		})

		stringDuplicateMapConcat := NewMapConcat(func(s string) []string {
			return []string{s, s}
		})

		intMapConcat.Connect(intToStringMapConcat)
		intToStringMapConcat.Connect(stringDuplicateMapConcat)

		expected := []string{"0", "0", "1", "1", "2", "2", "3", "3", "4", "4"}
		result := []string{}

		go func() {
			intMapConcat.In <- 5
			close(intMapConcat.In)
		}()

		go intMapConcat.Run()
		go intToStringMapConcat.Run()
		go stringDuplicateMapConcat.Run()

		for i := range stringDuplicateMapConcat.Out {
			result = append(result, i)
		}

		assert.ElementsMatch(t, expected, result)
	})

}
