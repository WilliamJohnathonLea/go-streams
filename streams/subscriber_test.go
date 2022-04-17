package streams

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceSubscriber(t *testing.T) {

	t.Run("slice subscriber attached to a slice publisher", func(t *testing.T) {
		expected := []int{1, 2, 3}
		pub := NewSlicePublisher(expected)
		sub := NewSliceSubscriber[int]()

		pub.Connect(sub)

		pub.Publish()
		sub.Subscribe()

		assert.ElementsMatch(t, expected, sub.Value())
	})

	t.Run("subscribe to publisher via a flow", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []string{"1", "2", "3"}
		pub := NewSlicePublisher(input)
		flow := NewMapConcat(func(i int) []string {
			return []string{fmt.Sprint(i)}
		})
		sub := NewSliceSubscriber[string]()

		pub.Connect(flow)
		flow.Connect(sub)

		pub.Publish()
		go flow.Run()
		sub.Subscribe()

		assert.ElementsMatch(t, expected, sub.Value())
	})

}
