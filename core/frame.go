package controller

import (
	"fmt"
	"log"
	"sync"
)

// type FrameMsg interface {
// 	InMsg(msg interface{})
// }

type Frame struct {
	ID          int
	Initialised bool
	Running     bool
	Cmd         chan int8
	sync.Mutex
	In chan interface{}
	// FnMap map[int8]interface{}
	Fn interface{}
}

// type Fnx struct {
// 	Fn interface{}
// }

// func (f *Frame) InMsg(msg interface{}) {
// 	log.Println("frame id:", f.ID, "chan:", f.In, "msg:", msg)
// }

func PrintFrame(f *Frame) {
	fmt.Println("frame id:", f.ID)
	fmt.Println("initialised:", f.Initialised)
	fmt.Println("running:", f.Running)
	fmt.Println("cmd chan:", f.Cmd)
	fmt.Println("in chan:", f.In)
	// fmt.Println("fn map:", f.FnMap)
	fmt.Println("fn:", f.Fn)
	// fmt.Println()
}

func NewFrame(id int) *Frame {
	return &Frame{
		ID:          id,
		Initialised: false,
		Running:     false,
		Cmd:         make(chan int8),
		In:          make(chan interface{}),
		// FnMap:       make(map[int8]interface{}),
	}
}

func (f *Frame) AddFn(id int8, fn interface{}) {
	f.Lock()
	// f.FnMap[id] = fn
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
		f.Initialised = true
		f.Unlock()
		for {
		loop:
			cmd := <-f.Cmd
			switch {
			case cmd == RUN:
				f.Running = true
				log.Printf("frame id: %v cmd: RUN aka %v\n", f.ID, cmd)

			case cmd == STOP:
				f.Running = false
				log.Printf("frame id: %v cmd: STOP aka %v\n", f.ID, cmd)
				goto loop
			}

			for {
				select {
				case cmd := <-f.Cmd:
					// log.Printf("echo: %v cmd: %v\n", f.ID, cmd)
					switch {
					case cmd == EXIT:
						f.Initialised = false
						f.Running = false
						log.Printf("frame id: %v cmd: EXIT aka %v\n", f.ID, cmd)
						return

					case cmd == STOP:
						f.Running = false
						log.Printf("frame id: %v cmd: STOP aka %v\n", f.ID, cmd)
						goto loop
					}

				case msg := <-f.In:
					// log.Println("frame id:", f.ID, "chan:", f.In, "msg:", msg)
					// f.InMsg(msg)
					// f.FnMap[0].(func(interface{}))(msg)
					f.Fn.(func(interface{}))(msg)
				}
			}
		}
	}()
}
