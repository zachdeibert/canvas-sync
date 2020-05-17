package task

//	#include <stdio.h>
//	#include <stdlib.h>
//	#include <sys/select.h>
//	#include <sys/types.h>
//	#include <termios.h>
//	#include <unistd.h>
//
//	static void *disable_buffer() {
//		struct termios *term = (struct termios *) malloc(sizeof(struct termios));
//		tcgetattr(STDIN_FILENO, term);
//		term->c_lflag &= ~(ICANON | ECHO);
//		tcsetattr(STDIN_FILENO, TCSANOW, term);
//		return term;
//	}
//
//	static void enable_buffer(void *_term) {
//		struct termios *term = (struct termios *) _term;
//		term->c_lflag |= ICANON | ECHO;
//		tcsetattr(STDIN_FILENO, TCSADRAIN, term);
//		free(term);
//	}
import "C"

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	monitorNameFrac = 0.3
)

type progressUpdate struct {
	row      int
	progress int
}

type renderJob struct {
	header   [][]string
	footer   [][]string
	progress []progressUpdate
	all      bool
}

type monitoredTask struct {
	task         *Task
	lastProgress int
	indent       []bool
}

type monitorLayout struct {
	screenWidth       int
	screenHeight      int
	headerHeight      int
	footerHeight      int
	nameWidth         int
	progressStart     int
	progressWidth     int
	tasks             []monitoredTask
	screenSizeChanged bool
}

type interactiveMonitor struct {
	root           *Task
	header         *Section
	footer         *Section
	layout         monitorLayout
	ticker         *time.Ticker
	tickerCleanup  chan interface{}
	stdout         *bufio.Writer
	closed         bool
	pendingJob     renderJob
	mutex          *sync.Mutex
	rendererIdle   bool
	rendererNotify chan interface{}
}

func createInteractiveMonitor(root *Task) *interactiveMonitor {
	m := &interactiveMonitor{
		root:   root,
		header: CreateSection(),
		footer: CreateSection(),
		layout: monitorLayout{
			screenWidth:  80,
			screenHeight: 24,
		},
		tickerCleanup: make(chan interface{}),
		stdout:        bufio.NewWriterSize(os.Stdout, 81920),
		closed:        false,
		pendingJob: renderJob{
			header:   nil,
			footer:   nil,
			progress: []progressUpdate{},
			all:      true,
		},
		mutex:          &sync.Mutex{},
		rendererIdle:   true,
		rendererNotify: make(chan interface{}),
	}
	m.header.AddListener(func(_ *Section, text [][]string) {
		m.mutex.Lock()
		m.pendingJob.header = text
		m.mutex.Unlock()
		m.notifyRenderer()
	})
	m.footer.AddListener(func(_ *Section, text [][]string) {
		m.mutex.Lock()
		m.pendingJob.footer = text
		m.mutex.Unlock()
		m.notifyRenderer()
	})
	progressListener := func(t *Task, progress float32) {
		m.mutex.Lock()
		if !m.pendingJob.all {
			for row, mon := range m.layout.tasks {
				if mon.task == t {
					bar := m.calculateProgressBar(progress)
					if bar != mon.lastProgress {
						m.pendingJob.progress = append(m.pendingJob.progress, progressUpdate{
							row:      row,
							progress: bar,
						})
					}
					break
				}
			}
		}
		m.mutex.Unlock()
		m.notifyRenderer()
	}
	finishListener := func(t *Task) {
		m.mutex.Lock()
		m.pendingJob.all = true
		m.mutex.Unlock()
		m.notifyRenderer()
	}
	var childrenListener func(parent *Task, child *Task, children []*Task)
	childrenListener = func(parent *Task, child *Task, children []*Task) {
		child.AddProgressListener(progressListener)
		child.AddChildrenListener(childrenListener)
		child.AddFinishListener(finishListener)
		m.mutex.Lock()
		m.pendingJob.all = true
		m.mutex.Unlock()
		m.notifyRenderer()
	}
	m.registerListeners(root, progressListener, childrenListener, finishListener)
	m.detectScreenSize(false)
	fmt.Fprintf(m.stdout, "\033[?25l\033[%d;1H", m.layout.screenHeight)
	for i := 1; i < m.layout.screenHeight; i++ {
		fmt.Fprintln(m.stdout)
	}
	fmt.Fprint(m.stdout, "\033[2J")
	m.ticker = time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case _ = <-m.ticker.C:
				m.detectScreenSize(true)
				break
			case _ = <-m.tickerCleanup:
				return
			}
		}
	}()
	go func() {
		for !m.closed {
			<-m.rendererNotify
			m.rendererIdle = false
			for !m.closed {
				m.mutex.Lock()
				job := m.pendingJob
				m.pendingJob = renderJob{
					header:   nil,
					footer:   nil,
					progress: []progressUpdate{},
					all:      false,
				}
				m.mutex.Unlock()
				if job.header == nil && job.footer == nil && len(job.progress) == 0 && !job.all {
					m.rendererIdle = true
					break
				}
				if job.all {
					m.relayout()
				} else {
					for _, p := range job.progress {
						m.renderProgress(p.row, p.progress)
					}
					if job.header != nil {
						m.renderHeader(job.header)
					}
					if job.footer != nil {
						m.renderFooter(job.footer)
					}
				}
				m.flushRender()
			}
		}
	}()
	m.notifyRenderer()
	return m
}

func (m *interactiveMonitor) notifyRenderer() {
	if m.rendererIdle {
		m.rendererNotify <- nil
	}
}

func (m *interactiveMonitor) registerListeners(t *Task, progressListener func(*Task, float32), childrenListener func(*Task, *Task, []*Task), finishListener func(*Task)) {
	t.AddProgressListener(progressListener)
	t.AddChildrenListener(childrenListener)
	t.AddFinishListener(finishListener)
	for _, child := range t.children {
		m.registerListeners(child, progressListener, childrenListener, finishListener)
	}
}

func (m *interactiveMonitor) Close() {
	if !m.closed {
		m.closed = true
		m.ticker.Stop()
		m.tickerCleanup <- nil
		m.rendererIdle = true
		m.notifyRenderer()
		fmt.Print("\033[1;1H\033[2J\033[?25h")
	}
}

func (m *interactiveMonitor) GetHeader() *Section {
	return m.header
}

func (m *interactiveMonitor) GetFooter() *Section {
	return m.footer
}

func (m *interactiveMonitor) IsInteractive() bool {
	return true
}

func (m *interactiveMonitor) flushRender() {
	fmt.Fprint(m.stdout, "\033[1;1H")
	if !m.closed {
		m.stdout.Flush()
	}
}

func (m *interactiveMonitor) renderSection(text [][]string, top int) {
	if m.layout.screenSizeChanged {
		m.relayout()
		return
	}
	for row, cols := range text {
		leftStr := cols[AlignLeft]
		centerStr := cols[AlignCenter]
		rightStr := cols[AlignRight]
		left := len(leftStr)
		center := len(centerStr)
		right := len(rightStr)
		if left+center+right+2 > m.layout.screenWidth {
			diff := (m.layout.screenWidth - (left + center + right + 2))
			diff = (diff / 3) * 3
			left -= diff
			center -= diff
			right -= diff
			leftStr = leftStr[0:left]
			centerStr = leftStr[0:center]
			rightStr = rightStr[0:right]
		}
		centerStart := (m.layout.screenWidth - center) / 2
		if centerStart < left+1 {
			centerStart = left + 1
		}
		fmt.Fprintf(m.stdout, fmt.Sprintf("\033[%%d;1H%%-%ds %%s %%%ds", centerStart-1, m.layout.screenWidth-(centerStart+center+2)), row+top, leftStr, centerStr, rightStr)
	}
}

func (m *interactiveMonitor) renderHeader(text [][]string) {
	if m.layout.screenSizeChanged {
		m.relayout()
		return
	}
	if len(text) != m.layout.headerHeight {
		m.relayout()
	} else {
		m.renderSection(text, 1)
	}
}

func (m *interactiveMonitor) renderFooter(text [][]string) {
	if m.layout.screenSizeChanged {
		m.relayout()
		return
	}
	if len(text) != m.layout.footerHeight {
		m.relayout()
	} else {
		m.renderSection(text, m.layout.screenHeight-m.layout.footerHeight)
	}
}

func (m *interactiveMonitor) renderProgress(row int, progress int) {
	if m.layout.screenSizeChanged {
		m.relayout()
		return
	}
	if m.layout.tasks[row].lastProgress < 0 {
		m.layout.tasks[row].lastProgress = 0
	}
	start := m.layout.progressStart + m.layout.tasks[row].lastProgress
	end := m.layout.progressStart + progress
	m.layout.tasks[row].lastProgress = progress
	var r string
	if start < end {
		r = "#"
	} else {
		r = " "
		tmp := start
		start = end
		end = tmp
	}
	fmt.Fprintf(m.stdout, "\033[%d;%dH", 1+m.layout.headerHeight+row, 1+start)
	for i := start; i < end; i++ {
		fmt.Fprint(m.stdout, r)
	}
}

func (m *interactiveMonitor) calculateProgressBar(progress float32) int {
	if m.layout.progressWidth <= 0 {
		return -1
	}
	return int(progress * float32(m.layout.progressWidth))
}

func (m *interactiveMonitor) renderAll() {
	if m.layout.screenSizeChanged {
		m.relayout()
		return
	}
	fmt.Fprint(m.stdout, "\033[2J")
	m.renderHeader(m.header.GetText())
	m.renderFooter(m.footer.GetText())
	for i, mon := range m.layout.tasks {
		fmt.Fprintf(m.stdout, "\033[%d;1H\033[2K", 1+m.layout.headerHeight+i)
		for i, v := range mon.indent {
			if v {
				if i == len(mon.indent)-1 {
					fmt.Fprint(m.stdout, " \u2523\u2501")
				} else {
					fmt.Fprint(m.stdout, " \u2503 ")
				}
			} else {
				if i == len(mon.indent)-1 {
					fmt.Fprint(m.stdout, " \u2517\u2501")
				} else {
					fmt.Fprint(m.stdout, "   ")
				}
			}
		}
		name := mon.task.name
		max := m.layout.nameWidth - 3*len(mon.indent)
		if len(name) > max {
			name = name[0:max]
		}
		fmt.Fprintf(m.stdout, fmt.Sprintf("%%-%ds [\033[%%dG]", max), name, m.layout.screenWidth-1)
		m.layout.tasks[i].lastProgress = -1
		m.renderProgress(i, m.calculateProgressBar(mon.task.GetProgress()))
	}
}

func (m *interactiveMonitor) readScreenSize() (int, int) {
	term := C.disable_buffer()
	defer C.enable_buffer(term)
	os.Stdout.WriteString("\033[10000;10000H\033[6n")
	var rows int
	var cols int
	fmt.Scanf("\033[%d;%dR", &rows, &cols)
	return rows, cols
}

func (m *interactiveMonitor) detectScreenSize(locking bool) {
	rows, cols := m.readScreenSize()
	if (m.layout.screenWidth != cols || m.layout.screenHeight != rows) && cols != 0 && rows != 0 {
		if locking {
			m.mutex.Lock()
			defer m.mutex.Unlock()
		}
		m.layout.screenWidth = cols
		m.layout.screenHeight = rows
		m.layout.screenSizeChanged = true
	}
}

func (m *interactiveMonitor) layoutTask(task *Task, layout []monitoredTask, usedTasks int) int {
	i := usedTasks
	for _, child := range task.children {
		used := m.layoutTask(child, layout, i+1)
		if (used > 0 || child.state != taskStateFinished) && i < len(layout) {
			layout[i] = monitoredTask{
				task:         child,
				lastProgress: 0,
			}
			i += used + 1
		}
	}
	return i - usedTasks
}

func containsTask(list []*Task, search *Task) bool {
	for _, t := range list {
		if t == search {
			return true
		}
	}
	return false
}

func (m *interactiveMonitor) cleanupIndent(tasks []monitoredTask, before int, iter int) {
	idx := len(tasks[before-1].indent) - 1 - iter
	for tid := before - 1; tid >= 0; tid-- {
		tasks[tid].indent[idx] = false
		if len(tasks[tid].indent) == idx+1 {
			return
		}
	}
}

func (m *interactiveMonitor) relayout() {
	layout := monitorLayout{
		screenWidth:  m.layout.screenWidth,
		screenHeight: m.layout.screenHeight,
		headerHeight: len(m.header.GetText()),
		footerHeight: len(m.footer.GetText()),
		nameWidth:    int(float32(m.layout.screenWidth-1) * monitorNameFrac),
	}
	layout.progressStart = layout.nameWidth + 2
	layout.progressWidth = layout.screenWidth - layout.progressStart - 2
	layout.tasks = make([]monitoredTask, layout.screenHeight-layout.headerHeight-layout.footerHeight)
	used := m.layoutTask(m.root, layout.tasks, 0)
	layout.tasks = layout.tasks[0:used]
	indentTemplate := []bool{}
	last := m.root
	for i, t := range layout.tasks {
		var indent []bool
		if containsTask(last.children, t.task) {
			indent = make([]bool, len(indentTemplate))
			copy(indent, indentTemplate)
		} else if containsTask(last.children, t.task.parent) {
			last = t.task.parent
			indent = make([]bool, len(indentTemplate)+1)
			copy(indent, indentTemplate)
			indent[len(indentTemplate)] = true
			indentTemplate = make([]bool, len(indent))
			copy(indentTemplate, indent)
		} else {
			c := 0
			for !containsTask(last.children, t.task) {
				last = last.parent
				indentTemplate = indentTemplate[0 : len(indentTemplate)-1]
				m.cleanupIndent(layout.tasks, i, c)
				c++
			}
			indent = make([]bool, len(indentTemplate))
			copy(indent, indentTemplate)
		}
		layout.tasks[i].indent = indent
	}
	for i := range indentTemplate {
		m.cleanupIndent(layout.tasks, len(layout.tasks), i)
	}
	m.layout = layout
	m.renderAll()
}
