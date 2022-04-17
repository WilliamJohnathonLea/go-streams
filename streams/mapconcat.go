package streams

type MapConcat[In, Out any] struct {
	in  chan In
	out chan Out
	fun func(In) []Out
}

func NewMapConcat[In, Out any](f func(In) []Out) *MapConcat[In, Out] {
	return &MapConcat[In, Out]{
		in:  make(chan In),
		out: make(chan Out),
		fun: f,
	}
}

func (mc *MapConcat[In, Out]) Inlet() chan In {
	return mc.in
}

func (mc *MapConcat[In, Out]) Connect(o Inletter[Out]) {
	mc.out = o.Inlet()
}

func (mc *MapConcat[In, Out]) Execute() {
	go func() {
		defer close(mc.out)
		outData := []Out{}

		for i := range mc.in {
			outData = append(outData, mc.fun(i)...)
		}

		for _, o := range outData {
			mc.out <- o
		}
	}()
}
