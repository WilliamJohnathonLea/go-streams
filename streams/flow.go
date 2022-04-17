package streams

type Flow[In, Out any] interface {
	Inletter[In]
	StreamConnector[Out]
	StreamExector
	Flow()
}

func MapConcatFlow[In, Out any](f func(In) []Out) Flow[In, Out] {
	return newMapConcat(f)
}

func MapFlow[In, Out any](f func(In) Out) Flow[In, Out] {
	return newMap(f)
}
