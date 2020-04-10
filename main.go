package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"./task"
)

func main() {
	/*
		subdomain := ""
		if len(os.Args) >= 2 {
			subdomain = os.Args[1]
		}
		token := ""
		if len(os.Args) == 2 {
			var err error
			var b []byte
			if b, err = ioutil.ReadFile(fmt.Sprintf("%s.pri", subdomain)); err != nil {
				if os.IsNotExist(err) {
					fmt.Fprintf(os.Stderr, "Error: no authentication token specified but file '%s.pri' does not exist.\n", subdomain)
				}
			}
			token = string(b)
		} else {
			token = os.Args[2]
		}
		if len(subdomain) == 0 || len(token) == 0 || len(os.Args) > 3 {
			fmt.Fprintf(os.Stderr, "Usage: %s <canvas subdomain> [authenication token]\n"+
				"\n"+
				"canvas subdomain:    This is the subdomain of instructure.com to use.\n"+
				"                     For example, this would be 'canvas' for the domain 'canvas.instructure.com'.\n"+
				"\n"+
				"authenication token: This is the authentication token to use for connecting to Canvas.\n"+
				"                     If this argument is not given, a file named <canvas subdomain>.pri must be\n"+
				"                     present in the current directory that contains the token.\n", os.Args[0])
			os.Exit(1)
		}
		c, err := canvas.CreateCanvas(strings.TrimSpace(subdomain), strings.TrimSpace(token))
		if err != nil {
			panic(err)
		}
		if err = c.Request("courses", nil, &[]model.Course{}, func(obj interface{}) error {
			courses := obj.(*[]model.Course)
			fmt.Println(courses)
			return nil
		}); err != nil {
			panic(err)
		}
	*/
	root := task.CreateRootTask()
	a := root.CreateSubtask("Task A", nil)
	a.CreateSubtask("Task A-1", nil)
	a2 := a.CreateSubtask("Task A-2", nil)
	a2.CreateSubtask("Task A-2-1", nil)
	a2.CreateSubtask("Task A-2-2", nil)
	a2.CreateSubtask("Task A-2-3", nil)
	a.CreateSubtask("Task A-3", nil)
	a4 := a.CreateSubtask("Task A-4", nil)
	a4.CreateSubtask("Task A-4-1", nil)
	a42 := a4.CreateSubtask("Task A-4-2", nil)
	a42.CreateSubtask("Task A-4-2-1", nil)
	a43 := a4.CreateSubtask("Task A-4-3", func(_ *task.Task, finish func()) {
		time.Sleep(5 * time.Second)
		finish()
	})
	b := root.CreateSubtask("Task B", nil)
	b1 := b.CreateSubtask("Task B-1", nil)
	b1.CreateSubtask("Task B-1-1", nil)
	mon := task.CreateMonitor(root)
	header := mon.GetHeader()
	header.SetSize(2)
	header.SetText(0, task.AlignLeft, "Left Text")
	header.SetText(0, task.AlignCenter, "Center Text")
	header.SetText(0, task.AlignRight, "Right Text")
	p := a43.CreateProgress(1)
	p.SetWork(50)
	t := time.Tick(time.Second / 10)
	defer mon.Close()
	c := make(chan os.Signal)
	a43.Start()
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-c:
			return
		case <-t:
			p.Finish(1)
			break
		}
	}
}
