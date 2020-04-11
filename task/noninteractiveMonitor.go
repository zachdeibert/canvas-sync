package task

import (
	"fmt"
	"strings"
)

type noninteractiveMonitor struct {
	header *Section
	footer *Section
}

func createNoninteractiveMonitor(root *Task) *noninteractiveMonitor {
	m := &noninteractiveMonitor{
		header: CreateSection(),
		footer: CreateSection(),
	}
	finishListener := func(t *Task) {
		names := []string{}
		for task := t; task != root; task = task.parent {
			names = append([]string{fmt.Sprintf("'%s'", task.name)}, names...)
		}
		if len(names) > 0 {
			fmt.Printf("Task %s done\n", strings.Join(names, " > "))
		}
	}
	var childrenListener func(*Task, *Task, []*Task)
	childrenListener = func(parent *Task, child *Task, children []*Task) {
		child.AddFinishListener(finishListener)
		child.AddChildrenListener(childrenListener)
	}
	m.registerListeners(root, finishListener, childrenListener)
	return m
}

func (m *noninteractiveMonitor) registerListeners(t *Task, finishListener func(*Task), childrenListener func(*Task, *Task, []*Task)) {
	t.AddFinishListener(finishListener)
	t.AddChildrenListener(childrenListener)
	for _, child := range t.children {
		m.registerListeners(child, finishListener, childrenListener)
	}
}

func (m *noninteractiveMonitor) Close() {
}

func (m *noninteractiveMonitor) GetHeader() *Section {
	return m.header
}

func (m *noninteractiveMonitor) GetFooter() *Section {
	return m.footer
}
