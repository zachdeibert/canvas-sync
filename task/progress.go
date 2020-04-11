package task

// Progress represents the progress of a task
type Progress struct {
	total        int
	done         int
	listeners    []func(*Progress, float32)
	lastDispatch float32
}

// CreateProgress creates a new Progress object
func CreateProgress() *Progress {
	return &Progress{
		total:        0,
		done:         0,
		listeners:    []func(*Progress, float32){},
		lastDispatch: -1,
	}
}

// SetWork sets the total number of units of work to do
func (p *Progress) SetWork(totalUnits int) {
	p.total = totalUnits
	p.dispatch()
}

// AddWork adds more to the total number of units of work to do
func (p *Progress) AddWork(moreUnits int) {
	p.total += moreUnits
	p.dispatch()
}

// Finish a quantity of work units
func (p *Progress) Finish(units int) {
	if n := p.done + units; n > p.total {
		p.done = p.total
	} else {
		p.done = n
	}
	p.dispatch()
}

// GetStatus returns how much progress has been made
func (p *Progress) GetStatus() float32 {
	if p.total == 0 {
		return 0
	}
	return float32(p.done) / float32(p.total)
}

func (p *Progress) dispatch() {
	val := p.GetStatus()
	if val != p.lastDispatch {
		p.lastDispatch = val
		for _, l := range p.listeners {
			l(p, val)
		}
	}
}

// AddListener adds a new listener that's fired every time the progress changes
func (p *Progress) AddListener(listener func(*Progress, float32)) {
	p.listeners = append(p.listeners, listener)
}
