package streams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {

	t.Run("map is a flow", func(t *testing.T) {
		assert.Implements(t, (*Flow[int, int])(nil), new(Map[int, int]))
	})

	t.Run("singular flow", func(t *testing.T) {
		intMapConcat := newMap(func(i int) int {
			return i * 2
		})
		expected := []int{0, 2, 4, 6, 8}
		result := []int{}

		go func() {
			for i := 0; i < 5; i++ {
				intMapConcat.in <- i
			}
			close(intMapConcat.in)
		}()

		go intMapConcat.Flow()

		for i := range intMapConcat.out {
			result = append(result, i)
		}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("two connected flows", func(t *testing.T) {
		intMapConcat := newMap(func(i int) int {
			return i
		})
		intToStringMapConcat := newMap(func(i int) string {
			return fmt.Sprint(i)
		})

		intMapConcat.Connect(intToStringMapConcat)

		expected := []string{"0", "1", "2", "3", "4"}
		result := []string{}

		go func() {
			for i := 0; i < 5; i++ {
				intMapConcat.in <- i
			}
			close(intMapConcat.in)
		}()

		go intMapConcat.Flow()
		go intToStringMapConcat.Flow()

		for i := range intToStringMapConcat.out {
			result = append(result, i)
		}

		assert.ElementsMatch(t, expected, result)
	})

}
