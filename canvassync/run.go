package canvassync

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/zachdeibert/canvas-sync/canvas"
	"github.com/zachdeibert/canvas-sync/task"
)

type taskPanic struct {
	t   *task.Task
	err interface{}
}

// Run the Canvas Sync program
func Run(c *canvas.Canvas) {
	dummyRoot := task.CreateRootTask()
	dummyExit := make(chan interface{})
	root := dummyRoot.CreateSubtask("Dummy", func(t *task.Task, finish func()) {
		<-dummyExit
		finish()
	})
	panicCh := make(chan taskPanic)
	root.AddPanicListener(func(src *task.Task, err interface{}) {
		select {
		case panicCh <- taskPanic{
			t:   src,
			err: err,
		}:
			break
		default:
			fmt.Fprintln(os.Stderr, "Unable to send panic to pipe")
			panic(err)
		}
	})
	name := make(chan string)
	dbCh := make(chan string)
	root.CreateSubtask("Database Check", databaseCheckTask(c, name, dbCh))
	coursesCh := make(chan []courseDiscoveryResult)
	root.CreateSubtask("Course Discovery", courseDiscoveryTask(c, coursesCh))
	mon := task.CreateMonitor(root)
	defer mon.Close()
	header := mon.GetHeader()
	header.SetSize(3)
	header.SetText(0, task.AlignCenter, "Canvas Sync Utility")
	header.SetText(1, task.AlignCenter, c.GetBaseURL())
	manager := task.CreateManager(dummyRoot, []int{0, 1, 1, 3}, 0)
	ch := make(chan os.Signal)
	manager.AddListener(func() {
		ch <- nil
	})
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	var db string
	for {
		select {
		case <-ch:
			return
		case n := <-name:
			header.SetText(1, task.AlignLeft, n)
			break
		case db = <-dbCh:
			break
		case courses := <-coursesCh:
			for _, course := range courses {
				root.CreateSubtask(fmt.Sprintf("Sync '%s'", course.name), courseTaskGroup(c, db, course)).Start()
			}
			root.CreateSubtask("Write Database to Disk", writeDatabaseTask(db))
			dummyExit <- nil
			break
		case err := <-panicCh:
			mon.Close()
			fmt.Fprintf(os.Stderr, "Panic in task %s\n", err.t.GetName(root))
			panic(err.err)
		}
	}
}
