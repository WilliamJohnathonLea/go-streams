package streams

// Subscriber is the termination point of a stream.
//
// The type In is the type the Subscriber will receive from the stream.
// The type Out is the type the Subscriber will produce.
type Subscriber[In, Out any] interface {
	StreamExector
	Inletter[In]
	Subscribe()
	Value() Out
}

// SliceSubscriber is a Subscriber which produces a slice of type T.
type SliceSubscriber[T any] struct {
	data []T
	in   chan T
}

// NewSliceSubscriber creates a SliceSubscriber and exposes it as
// a Subscriber.
func NewSliceSubscriber[T any]() Subscriber[T, []T] {
	return &SliceSubscriber[T]{
		data: []T{},
		in:   make(chan T),
	}
}

// Inlet exposes the inbound channel of the SliceSubscriber.
func (ss *SliceSubscriber[T]) Inlet() chan T {
	return ss.in
}

// Subscribe reads the inbound channel and appends the data to
// the internal slice.
func (ss *SliceSubscriber[T]) Subscribe() {
	for x := range ss.in {
		ss.data = append(ss.data, x)
	}
}

// Execute calls Subscribe. It is used when a Stream call its Run method.
func (ss *SliceSubscriber[T]) Execute() {
	ss.Subscribe()
}

// Value returns the slice held in the SliceSubscriber.
func (ss *SliceSubscriber[T]) Value() []T {
	return ss.data
}
