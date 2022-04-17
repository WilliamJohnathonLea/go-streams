# Go Streams

[![unit-tests Actions Status](https://github.com/WilliamJohnathonLea/go-streams/workflows/unit-tests/badge.svg)](https://github.com/WilliamJohnathonLea/go-streams/actions)

Easily build streamed operations in Go.

## Example stream
```go
pub := NewSlicePublisher([]int{0, 1, 2})
flow := MapFlow(func(i int) string {
	return fmt.Sprint(i)
})
sub := NewSliceSubscriber[string]()

pub.Connect(flow)
flow.Connect(sub)

stream := New(pub, flow, sub)
stream.Run()
```

## Built in Publishers
- SlicePublisher
- more to come...
## Built in Subscribers
- SliceSubscriber
- more to come...
