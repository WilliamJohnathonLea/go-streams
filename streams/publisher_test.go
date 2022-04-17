package streams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicePublisher(t *testing.T) {

	expected := []int{0, 1, 2}
	publisher := NewSlicePublisher(expected)
	out := publisher.Publish()
	actual := []int{}

	for i := range out {
		actual = append(actual, i)
	}

	assert.ElementsMatch(t, expected, actual)
}
