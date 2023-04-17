package pair

type Pair[L any, R any] interface {
	L() L
	R() R
}

type immutablePair[L any, R any] struct {
	l L
	r R
}

func New[L any, R any](l L, r R) *immutablePair[L, R] {
	return &immutablePair[L, R]{
		l: l,
		r: r,
	}
}

func (p *immutablePair[L, R]) L() L {
	return p.l
}

func (p *immutablePair[L, R]) R() R {
	return p.r
}
