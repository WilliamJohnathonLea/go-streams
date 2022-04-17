package streams

type Publisher[Out any] interface {
	StreamExector
	StreamConnector[Out]
	Publish() <-chan Out
}

type SlicePublisher[Out any] struct {
	data []Out
	out  chan Out
}

func NewSlicePublisher[Out any](slice []Out) Publisher[Out] {
	return &SlicePublisher[Out]{
		data: slice,
		out:  make(chan Out),
	}
}

func (sp *SlicePublisher[Out]) Publish() <-chan Out {
	for _, d := range sp.data {
		sp.out <- d
	}
	close(sp.out)
	return sp.out
}

func (sp *SlicePublisher[Out]) Execute() {
	sp.Publish()
}

func (sp *SlicePublisher[Out]) Connect(o Inletter[Out]) {
	sp.out = o.Inlet()
}
