package streams

type Inletter[T any] interface {
	Inlet() chan T
}

type StreamConnector[T any] interface {
	Connect(Inletter[T])
}

type StreamExector interface {
	Execute()
}

type Stream []StreamExector

func New(xs ...StreamExector) (Stream) {
	return xs
}

func (s Stream) Run() {
	for _, i := range s {
		go func(se StreamExector) {
			se.Execute()
		}(i)
	}
}
