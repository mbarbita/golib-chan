package controller

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Frame struct {
	ID          int
	Initialised bool
	Running     bool
	Cmd         chan int8
	sync.Mutex
	In chan interface{}
	Fn interface{}
}

func PrintFrame(f *Frame) {
	fmt.Println("frame id:", f.ID)
	fmt.Println("initialised:", f.Initialised)
	fmt.Println("running:", f.Running)
	fmt.Println("cmd chan:", f.Cmd)
	fmt.Println("in chan:", f.In)
	fmt.Println("fn:", f.Fn)
}

func NewFrame(id int) *Frame {
	return &Frame{
		ID:          id,
		Initialised: false,
		Running:     false,
		Cmd:         make(chan int8),
		In:          make(chan interface{}),
	}
}

func (f *Frame) SetFn(fn interface{}) {
	f.Lock()
	f.Fn = fn
	f.Unlock()
}

func (f *Frame) Run() {
	f.Cmd <- RUN
}

// Stop ...
func (f *Frame) Stop() {
	f.Cmd <- STOP
}

func (f *Frame) Init() {
	go func() {
		f.Lock()
		fcast := f.Fn.(func(interface{}))
		f.Initialised = true
		f.Unlock()
		for {
		loop:
			cmd := <-f.Cmd
			switch {
			case cmd == RUN:
				f.Running = true
				log.Printf("frame id: %v cmd: %v aka RUN\n", f.ID, cmd)

			case cmd == STOP:
				f.Running = false
				log.Printf("frame id: %v cmd: %v aka STOP\n", f.ID, cmd)
				goto loop
			}

			for {
				select {
				case cmd := <-f.Cmd:
					switch {
					case cmd == EXIT:
						f.Lock()
						f.Initialised = false
						f.Running = false
						f.Unlock()
						log.Printf("frame id: %v cmd: %v aka EXIT\n", f.ID, cmd)
						return

					case cmd == STOP:
						f.Lock()
						f.Running = false
						f.Unlock()
						log.Printf("frame id: %v cmd: %v aka STOP\n", f.ID, cmd)
						goto loop
					}
				case msg := <-f.In:
					start := time.Now()
					fcast(msg)
					elapsed := time.Since(start)
					log.Println("frame id:", f.ID, "fn call duration:", elapsed)
				}
			}
		}
	}()
}
