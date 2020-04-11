package canvassync

import (
	"os"
	"os/signal"
	"syscall"

	"../canvas"
	"../task"
)

// Run the Canvas Sync program
func Run(c *canvas.Canvas) {
	root := task.CreateRootTask()
	name := make(chan string)
	dbCh := make(chan string)
	root.CreateSubtask("Database Check", databaseCheckTask(c, name, dbCh))
	mon := task.CreateMonitor(root)
	defer mon.Close()
	header := mon.GetHeader()
	header.SetSize(3)
	header.SetText(0, task.AlignCenter, "Canvas Sync Utility")
	header.SetText(1, task.AlignCenter, c.GetBaseURL())
	manager := task.CreateManager(root, []int{0, 1, 3}, 0)
	ch := make(chan os.Signal)
	manager.AddListener(func() {
		ch <- nil
	})
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-ch:
			return
		case n := <-name:
			header.SetText(1, task.AlignLeft, n)
			break
		case <-dbCh:
			break
		}
	}
}
