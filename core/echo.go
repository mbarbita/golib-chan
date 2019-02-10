package controller

import (
	"fmt"
	"log"
	"sync"
)

type Echo struct {
	ID      int
	Running bool
	Cmd     chan int8
	sync.Mutex
	In chan interface{}
}

func NewEcho(id int, cmd chan int8, inCh chan interface{}) *Echo {
	return &Echo{
		ID:      id,
		Running: false,
		Cmd:     cmd,
		In:      inCh,
	}
}

func PrintComp(e *Echo) {
	fmt.Println("Echo id:", e.ID)
	fmt.Println("running:", e.Running)
	fmt.Println("cmd chan:", e.Cmd)
	fmt.Println("in chan:", e.In)
	fmt.Println()
}

func (e *Echo) Start() {
	go func() {
		for {
			cmd := <-e.Cmd
			if cmd == 1 {
				e.Lock()
				e.Running = true
				e.Unlock()
			}

			if cmd == 0 {
				e.Lock()
				e.Running = false
				e.Unlock()
				return
			}

			for {
				if !e.Running {
					break
				}
				select {
				case cmd := <-e.Cmd:
					log.Printf("echo: %v cmd: %v\n", e.ID, cmd)
					if (cmd == 0) || (cmd == 2) {
						e.Running = false
						break
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
	e.Lock()
	e.Running = false
	e.Unlock()
}
