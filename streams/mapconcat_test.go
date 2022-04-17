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
			intMapConcat.in <- 5
			close(intMapConcat.in)
		}()

		intMapConcat.Execute()

		for i := range intMapConcat.out {
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
			intMapConcat.in <- 5
			close(intMapConcat.in)
		}()

		intMapConcat.Execute()
		intToStringMapConcat.Execute()

		for i := range intToStringMapConcat.out {
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
			intMapConcat.in <- 5
			close(intMapConcat.in)
		}()

		intMapConcat.Execute()
		intToStringMapConcat.Execute()
		stringDuplicateMapConcat.Execute()

		for i := range stringDuplicateMapConcat.out {
			result = append(result, i)
		}

		assert.ElementsMatch(t, expected, result)
	})

}
