package task

import "sync"

type calculationResult int

const (
	calcLevelNotEmpty calculationResult = (1 << iota)
	calcLevelFinished calculationResult = (1 << iota)
)

// Manager starts tasks when other tasks finish to keep system load somewhat constant
type Manager struct {
	root      *Task
	maxes     []int
	def       int
	listeners []func()
	mutex     *sync.Mutex
}

// CreateManager creates a new Manager
func CreateManager(root *Task, maxes []int, def int) *Manager {
	m := &Manager{
		root:  root,
		maxes: maxes,
		def:   def,
		mutex: &sync.Mutex{},
	}
	finishListener := func(t *Task) {
		m.mutex.Lock()
		defer m.mutex.Unlock()
		if l := m.getLevel(t); (m.recalculate(l) & calcLevelFinished) != 0 {
			for i := 0; ; i++ {
				if i != l {
					res := m.recalculate(i)
					if (res & calcLevelNotEmpty) == 0 {
						break
					}
					if (res & calcLevelFinished) == 0 {
						return
					}
				}
			}
			m.dispatch()
		}
	}
	var childrenListener func(*Task, *Task, []*Task)
	childrenListener = func(parent *Task, child *Task, children []*Task) {
		child.AddChildrenListener(childrenListener)
		child.AddFinishListener(finishListener)
		m.mutex.Lock()
		defer m.mutex.Unlock()
		m.recalculate(m.getLevel(parent) + 1)
	}
	m.registerListeners(root, childrenListener, finishListener)
	for i := 0; (m.recalculate(i) & calcLevelNotEmpty) != 0; i++ {
	}
	return m
}

func (m *Manager) registerListeners(t *Task, childrenListener func(*Task, *Task, []*Task), finishListener func(*Task)) {
	t.AddChildrenListener(childrenListener)
	t.AddFinishListener(finishListener)
	for _, child := range t.children {
		m.registerListeners(child, childrenListener, finishListener)
	}
}

func (m *Manager) getLevel(t *Task) int {
	level := 0
	for task := t; task != m.root; level++ {
		task = task.parent
	}
	return level
}

func (m *Manager) nextSibling(t *Task) *Task {
	i := 0
	if t.parent == nil {
		return nil
	}
	for ; i < len(t.parent.children)-1; i++ {
		if t.parent.children[i] == t {
			return t.parent.children[i+1]
		}
	}
	return nil
}

func (m *Manager) recalculate(level int) calculationResult {
	limit := m.def
	if level < len(m.maxes) {
		limit = m.maxes[level]
	}
	t := m.root
	l := 0
	levelNotEmpty := false
	levelFinished := true
	for t != nil {
		if l < level {
			if len(t.children) > 0 {
				t = t.children[0]
				l++
			} else {
				nt := m.nextSibling(t)
				for nt == nil && t.parent != nil {
					t = t.parent
					nt = m.nextSibling(t)
					l--
				}
				t = nt
			}
		} else {
			levelNotEmpty = true
			if t.state != taskStateFinished {
				levelFinished = false
			}
			if limit > 0 && t.state == taskStateQueued {
				t.Start()
				limit--
			}
			nt := m.nextSibling(t)
			for nt == nil && t.parent != nil {
				t = t.parent
				nt = m.nextSibling(t)
				l--
			}
			t = nt
		}
	}
	var res calculationResult = 0
	if levelNotEmpty {
		res |= calcLevelNotEmpty
	}
	if levelFinished {
		res |= calcLevelFinished
	}
	return res
}

func (m *Manager) dispatch() {
	for _, l := range m.listeners {
		l()
	}
}

// AddListener adds a new listener that's fired once all tasks are finished
func (m *Manager) AddListener(listener func()) {
	m.listeners = append(m.listeners, listener)
}
