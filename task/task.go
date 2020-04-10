package task

type taskState int

const (
	taskStateQueued   taskState = iota
	taskStateRunning  taskState = iota
	taskStateFinished taskState = iota
)

type taskProgress struct {
	scale    float32
	progress *Progress
}

// Task represents a group of work to do
type Task struct {
	name                 string
	parent               *Task
	children             []*Task
	progress             []taskProgress
	state                taskState
	startFunc            func(*Task, func())
	progressListeners    []func(*Task, float32)
	lastProgressDispatch float32
	childrenListeners    []func(*Task, *Task, []*Task)
	finishListeners      []func(*Task)
}

// CreateRootTask creates a root task
func CreateRootTask() *Task {
	return &Task{
		name:                 "RootTask",
		parent:               nil,
		children:             []*Task{},
		progress:             []taskProgress{},
		state:                taskStateQueued,
		startFunc:            nil,
		progressListeners:    []func(*Task, float32){},
		lastProgressDispatch: -1,
		childrenListeners:    []func(*Task, *Task, []*Task){},
	}
}

// Start starts the task
func (t *Task) Start() {
	if t.startFunc != nil && t.state == taskStateQueued {
		t.state = taskStateRunning
		go t.startFunc(t, func() {
			t.state = taskStateFinished
			t.dispatchFinish()
		})
	}
}

// CreateProgress creates a new progress tracker for the task
func (t *Task) CreateProgress(scale float32) *Progress {
	progress := CreateProgress()
	t.progress = append(t.progress, taskProgress{
		scale:    scale,
		progress: progress,
	})
	progress.AddListener(func(_ *Progress, val float32) {
		t.dispatchProgress()
	})
	return progress
}

// GetProgress for the task
func (t *Task) GetProgress() float32 {
	var sum float32 = 0
	var total float32 = 0
	for _, p := range t.progress {
		sum += p.scale * p.progress.GetStatus()
		total += p.scale
	}
	if total == 0 {
		return 0
	}
	return sum / total
}

// CreateSubtask creates a new task that runs under this task
func (t *Task) CreateSubtask(name string, start func(*Task, func())) *Task {
	task := &Task{
		name:                 name,
		parent:               t,
		children:             []*Task{},
		progress:             []taskProgress{},
		state:                taskStateQueued,
		startFunc:            start,
		progressListeners:    []func(*Task, float32){},
		lastProgressDispatch: -1,
		childrenListeners:    []func(*Task, *Task, []*Task){},
	}
	t.children = append(t.children, task)
	t.dispatchChildren(task)
	return task
}

func (t *Task) dispatchProgress() {
	val := t.GetProgress()
	if val != t.lastProgressDispatch {
		t.lastProgressDispatch = val
		for _, l := range t.progressListeners {
			l(t, val)
		}
	}
}

// AddProgressListener adds a new listener that's fired every time the progress changes
func (t *Task) AddProgressListener(listener func(*Task, float32)) {
	t.progressListeners = append(t.progressListeners, listener)
}

func (t *Task) dispatchChildren(new *Task) {
	for _, l := range t.childrenListeners {
		l(t, new, t.children)
	}
}

// AddChildrenListener adds a new listener that's fired every time a new child is added
func (t *Task) AddChildrenListener(listener func(*Task, *Task, []*Task)) {
	t.childrenListeners = append(t.childrenListeners, listener)
}

func (t *Task) dispatchFinish() {
	for _, l := range t.finishListeners {
		l(t)
	}
}

// AddFinishListener adds a new lister that's fired when the task finishes
func (t *Task) AddFinishListener(listener func(*Task)) {
	t.finishListeners = append(t.finishListeners, listener)
}
