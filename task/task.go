package task

import (
	"fmt"
	"strings"
)

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
	inheritingProgress   bool
	progressListeners    []func(*Task, float32)
	lastProgressDispatch float32
	childrenListeners    []func(*Task, *Task, []*Task)
	finishListeners      []func(*Task)
	panicListeners       []func(*Task, interface{})
}

// CreateRootTask creates a root task
func CreateRootTask() *Task {
	return &Task{
		name:                 "RootTask",
		parent:               nil,
		children:             []*Task{},
		progress:             []taskProgress{},
		state:                taskStateFinished,
		startFunc:            nil,
		inheritingProgress:   false,
		progressListeners:    []func(*Task, float32){},
		lastProgressDispatch: -1,
		childrenListeners:    []func(*Task, *Task, []*Task){},
		panicListeners:       []func(*Task, interface{}){},
	}
}

// Start starts the task
func (t *Task) Start() {
	if t.startFunc != nil && t.state == taskStateQueued {
		t.state = taskStateRunning
		go func() {
			defer func() {
				if err := recover(); err != nil {
					t.dispatchPanic(t, err)
				}
			}()
			t.startFunc(t, func() {
				t.state = taskStateFinished
				t.dispatchFinish()
			})
		}()
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

func (t *Task) getProgress() (float32, float32) {
	var sum float32 = 0
	var total float32 = 0
	for _, p := range t.progress {
		sum += p.scale * p.progress.GetStatus()
		total += p.scale
	}
	if t.inheritingProgress {
		for _, c := range t.children {
			ds, dt := c.getProgress()
			sum += ds
			total += dt
		}
	}
	return sum, total
}

// GetProgress for the task
func (t *Task) GetProgress() float32 {
	sum, total := t.getProgress()
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
	task.AddPanicListener(t.dispatchPanic)
	if t.inheritingProgress {
		t.addInheritProgressListener(task)
		t.dispatchProgress()
	}
	return task
}

// GetName gets the qualified name of the task, optionally up to a root task
func (t *Task) GetName(root *Task) string {
	parts := []string{
		fmt.Sprintf("'%s'", t.name),
	}
	if root != nil {
		t = t.parent
		for t != nil && t != root {
			parts = append([]string{
				fmt.Sprintf("'%s'", t.name),
			}, parts...)
		}
	}
	return strings.Join(parts, " > ")
}

func (t *Task) addInheritProgressListener(child *Task) {
	child.AddProgressListener(func(_ *Task, p float32) {
		t.dispatchProgress()
	})
}

// InheritProgress causes this task's progress to be based on its childrens' progress
func (t *Task) InheritProgress() {
	if !t.inheritingProgress {
		t.inheritingProgress = true
		for _, child := range t.children {
			t.addInheritProgressListener(child)
		}
	}
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

// AddFinishListener adds a new listener that's fired when the task finishes
func (t *Task) AddFinishListener(listener func(*Task)) {
	t.finishListeners = append(t.finishListeners, listener)
}

func (t *Task) dispatchPanic(src *Task, err interface{}) {
	for _, l := range t.panicListeners {
		l(src, err)
	}
}

// AddPanicListener adds a new listener that's fired when the task causes a panic
func (t *Task) AddPanicListener(listener func(*Task, interface{})) {
	t.panicListeners = append(t.panicListeners, listener)
}
