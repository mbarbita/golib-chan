package controller

import (
	"fmt"
)

type Echo struct {
	ID          int
	Initialised bool
	Running     bool
	Cmd         chan int8
	// sync.Mutex
	In chan interface{}
}

func NewEcho(id int, cmd chan int8, inCh chan interface{}) *Echo {
	return &Echo{
		ID:          id,
		Initialised: false,
		Running:     false,
		Cmd:         cmd,
		In:          inCh,
	}
}

func PrintComp(e *Echo) {
	fmt.Println("Echo id:", e.ID)
	fmt.Println("Initialised:", e.Initialised)
	fmt.Println("running:", e.Running)
	fmt.Println("cmd chan:", e.Cmd)
	fmt.Println("in chan:", e.In)
	fmt.Println()
}

func (e *Echo) Init() {
	go func() {
		e.Initialised = true
		for {
		loop:
			cmd := <-e.Cmd
			if cmd == RUN {
				e.Running = true
			}

			if cmd == STOP {
				e.Running = false
				goto loop
			}

			for {
				select {
				case cmd := <-e.Cmd:
					fmt.Printf("echo: %v cmd: %v\n", e.ID, cmd)
					if cmd == EXIT {
						e.Initialised = false
						e.Running = false
						return
					}

					if cmd == STOP {
						e.Running = false
						goto loop
					}

				case msg := <-e.In:
					// Print the message
					fmt.Println("echo id:", e.ID, "chan:", e.In, "msg:", msg)
				}
			}
		}
	}()
}

// Stop ...
func (e *Echo) Stop() {
	e.Cmd <- STOP
}
