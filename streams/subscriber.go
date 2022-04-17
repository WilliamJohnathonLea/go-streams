package streams

type Subscriber[In, Out any] interface {
	Inletter[In]
	Subscribe()
	Value() Out
}

type SliceSubscriber[T any] struct {
	data []T
	in   chan T
}

func NewSliceSubscriber[T any]() Subscriber[T, []T] {
	return &SliceSubscriber[T]{
		data: []T{},
		in:   make(chan T),
	}
}

func (ss *SliceSubscriber[T]) Inlet() chan T {
	return ss.in
}

func (ss *SliceSubscriber[T]) Subscribe() {
	for x := range ss.in {
		ss.data = append(ss.data, x)
	}
}

func (ss *SliceSubscriber[T]) Value() []T {
	return ss.data
}
