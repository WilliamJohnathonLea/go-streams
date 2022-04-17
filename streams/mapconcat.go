package streams

type MapConcat[In, Out any] struct {
	In   chan In
	Out  chan Out
	Func func(In) []Out
}

func NewMapConcat[In, Out any](f func(In) []Out) *MapConcat[In, Out] {
	return &MapConcat[In, Out]{
		In: make(chan In),
		Out: make(chan Out),
		Func: f,
	}
}

func (mc *MapConcat[In, Out]) Inlet() chan In {
	return mc.In
}

func (mc *MapConcat[In, Out]) Connect(o Inletter[Out]) {
	mc.Out = o.Inlet()
}

func (mc *MapConcat[In, Out]) Run() {
	defer close(mc.Out)
	outData := []Out{}

	for i := range mc.In {
		outData = append(outData, mc.Func(i)...)
	}

	for _, o := range outData {
		mc.Out <- o
	}
}
