package domain

type RowState int

const (
	Busy RowState = iota
	Free
)

type Row interface {
	Acquire() bool
	Release()
	State() RowState
	Info() string
}

type rowImpl struct {
	busyChan chan struct{}
}

func NewRow() *rowImpl {
	row := rowImpl{
		busyChan: make(chan struct{}, 1),
	}
	row.busyChan <- struct{}{}
	return &row
}

func (r *rowImpl) Release() {
	r.busyChan <- struct{}{}
}

func (r *rowImpl) Acquire() bool {
	// r.busyChan <- struct{}{}
	select {
	case <-r.busyChan:
		return true
	default:
		return false
	}
}

func (r *rowImpl) State() RowState {
	select {
	case <-r.busyChan:
		r.busyChan <- struct{}{}
		return Free
	default:
		return Busy
	}
}

func (r *rowImpl) Info() string {
	state := r.State()
	switch state {
	case Free:
		return "free"
	case Busy:
		return "busy"
	default:
		return ""
	}
}
