package streams

type Inletter[T any] interface {
	Inlet() chan T
}

type StreamConnector[T any] interface {
	Connect(Inletter[T])
}

