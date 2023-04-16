package pair

type Pair[L any, R any] struct {
	l L
	r R
}

func New[L any, R any](l L, r R) *Pair[L, R] {
	return &Pair[L, R]{
		l: l,
		r: r,
	}
}

func (p *Pair[L, R]) Left() L {
	return p.l
}

func (p *Pair[L, R]) Right() R {
	return p.r
}
