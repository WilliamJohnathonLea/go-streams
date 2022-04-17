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

	t.Run("slice publisher is a stream executor", func(t *testing.T) {
		assert.Implements(t, (*StreamExector)(nil), new(SlicePublisher[int]))
	})

	t.Run("slice publisher publishes data", func(t *testing.T) {
		expected := []int{0, 1, 2}
		outChan := make(chan int)
		publisher := &SlicePublisher[int]{
			data: expected,
			out: outChan,
		}
		go publisher.Publish()
		actual := []int{}

		for i := range outChan {
			actual = append(actual, i)
		}

		assert.ElementsMatch(t, expected, actual)
	})

}
