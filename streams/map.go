package streams

type Map[In, Out any] struct {
	in  chan In
	out chan Out
	fun func(In) Out
}

func newMap[In, Out any](f func(In) Out) *Map[In, Out] {
	return &Map[In, Out]{
		in:  make(chan In),
		out: make(chan Out),
		fun: f,
	}
}

func (mc *Map[In, Out]) Inlet() chan In {
	return mc.in
}

func (mc *Map[In, Out]) Connect(o Inletter[Out]) {
	mc.out = o.Inlet()
}

func (mc *Map[In, Out]) Flow() {
	go func() {
		defer close(mc.out)
		outData := []Out{}

		for i := range mc.in {
			outData = append(outData, mc.fun(i))
		}

		for _, o := range outData {
			mc.out <- o
		}
	}()
}
