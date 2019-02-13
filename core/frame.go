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
	In    chan interface{}
	Fn    interface{}
	DurCh chan time.Duration
}

func PrintFrame(f *Frame) {
	fmt.Printf("frame id: %v type: %T\n", f.ID, f.ID)
	fmt.Printf("initialised: %v type: %T\n", f.Initialised, f.Initialised)
	fmt.Printf("running: %v type: %T\n", f.Running, f.Running)
	fmt.Printf("cmd chan: %v type: %T\n", f.Cmd, f.Cmd)
	fmt.Printf("in chan: %v type: %T\n", f.In, f.In)
	fmt.Printf("fn: %v type: %T\n", f.Fn, f.Fn)
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

// func (f *Frame) SetFn(fn interface{}) {
// 	f.Lock()
// 	f.Fn = fn
// 	f.Unlock()
// }

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
		// fcast := f.Fn.(func(interface{}))
		f.Initialised = true
		f.Unlock()
		for {
		loop:
			cmd := <-f.Cmd
			switch cmd {
			case RUN:
				f.Running = true
				log.Printf("frame id: %v cmd: %v aka RUN\n", f.ID, cmd)

			case STOP:
				f.Running = false
				log.Printf("frame id: %v cmd: %v aka STOP\n", f.ID, cmd)
				goto loop
			}

			for {
				select {
				case cmd := <-f.Cmd:
					switch cmd {
					case EXIT:
						f.Lock()
						f.Initialised = false
						f.Running = false
						f.Unlock()
						log.Printf("frame id: %v cmd: %v aka EXIT\n", f.ID, cmd)
						return

					case STOP:
						f.Lock()
						f.Running = false
						f.Unlock()
						log.Printf("frame id: %v cmd: %v aka STOP\n", f.ID, cmd)
						goto loop
					}
				case msg := <-f.In:
					start := time.Now()
					fcast(msg)
					select {
					case f.DurCh <- time.Since(start):
					default:
					}
					// log.Println("frame id:", f.ID, "fn call duration:", elapsed)
				}
			}
		}
	}()
}
