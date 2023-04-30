package abort

type notification struct{}

type Signal <-chan notification
type Aborter chan notification

func New() Aborter {
	return make(Aborter)
}

func (a Aborter) Signal() Signal {
	return chan notification(a)
}

func (a Aborter) Abort() {
	close(a)
}
