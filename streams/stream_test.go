package streams

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {

	t.Run("stream of a publisher to a subscriber", func(t *testing.T) {
		inputs := []int{0, 1, 2}
		pub := NewSlicePublisher(inputs)
		sub := NewSliceSubscriber[int]()

		pub.Connect(sub)

		stream := New(pub, sub)
		stream.Run()
		// stream.Run is aysnchronous so we need to wait
		time.Sleep(time.Millisecond * 500)
		assert.ElementsMatch(t, inputs, sub.Value())
	})

	t.Run("stream of a publisher to a subscriber via a flow", func(t *testing.T) {
		inputs := []int{0, 1, 2}
		expected := []string{"0", "1", "2"}
		pub := NewSlicePublisher(inputs)
		flow := MapFlow(func(i int) string {
			return fmt.Sprint(i)
		})
		sub := NewSliceSubscriber[string]()

		pub.Connect(flow)
		flow.Connect(sub)

		stream := New(pub, flow, sub)
		stream.Run()
		// stream.Run is aysnchronous so we need to wait
		time.Sleep(time.Millisecond * 500)
		assert.ElementsMatch(t, expected, sub.Value())
	})

}
