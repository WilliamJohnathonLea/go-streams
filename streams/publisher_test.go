package streams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicePublisher(t *testing.T) {

	t.Run("slice publisher is a publisher", func(t *testing.T) {
		assert.Implements(t, (*Publisher[int])(nil), new(SlicePublisher[int]))
	})

	t.Run("slice publisher is a stream connector", func(t *testing.T) {
		assert.Implements(t, (*StreamConnector[int])(nil), new(SlicePublisher[int]))
	})

	t.Run("slice publisher publishes data", func(t *testing.T) {
		expected := []int{0, 1, 2}
		publisher := NewSlicePublisher(expected)
		out := publisher.Publish()
		actual := []int{}

		for i := range out {
			actual = append(actual, i)
		}

		assert.ElementsMatch(t, expected, actual)
	})

}
