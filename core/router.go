package controller

import (
	"fmt"
	"log"
)

// Router ...
type Router struct {
	Frame
	OutMap map[int]chan interface{} // out id = chan
}

// NewRouter ...
func NewRouter(id int, cmd chan int8, inCh chan interface{}) *Router {
	return &Router{
		Frame: Frame{
			ID:          id, //id
			Initialised: false,
			Running:     false, // running
			Cmd:         cmd,
			In:          inCh, //in
		},
		OutMap: make(map[int]chan interface{}), //out
	}
}

// PrintRouter ...
func PrintRouter(r *Router) {
	fmt.Println("Router id:", r.ID)
	fmt.Println("running  :", r.Running)
	fmt.Println("in chan  :", r.In)
	for k, v := range r.OutMap {
		fmt.Printf("out id : %v, chan: %v\n", k, v)
	}
	fmt.Println()
}

// ModifyOut ...
func (r *Router) ModOut(outID int, ch chan interface{}) {
	r.Lock()
	defer r.Unlock()
	r.OutMap[outID] = ch
}

func (r *Router) DelOut(outID int, ch chan interface{}) {
	r.Lock()
	defer r.Unlock()
	delete(r.OutMap, outID)
	close(ch)
}

// ModifyIn ...
func (r *Router) ModIn(ch chan interface{}) {
	r.Lock()
	defer r.Unlock()
	r.In = ch
	close(r.In)
}

// ChangeID ...
func (r *Router) ModID(id int) {
	r.Lock()
	defer r.Unlock()
	r.ID = id

}

// Start ...
func (r *Router) Start() {
	go func() {
		r.Lock()
		r.Running = true
		r.Unlock()
		for {
			if !r.Running {
				// return
				break
			}
			msg := <-r.In
			// loop trough client map and send the message
			for k, v := range r.OutMap {
				select {
				case v <- msg:
				default:
					log.Printf("router %v could not send to chan id: %v chan:%v\n",
						r.ID, k, v)
				}
			}
		}
		// break
		log.Printf("router %v stopped.\n", r.ID)
	}()
}

// func (e *Echo) Init() {
// 	go func() {
// 		e.Lock()
// 		e.Initialised = true
// 		e.Unlock()
// 		for {
// 		loop:
// 			cmd := <-e.Cmd
// 			if cmd == RUN {
// 				e.Running = true
// 				log.Printf("echo id: %v cmd: RUN aka %v\n", e.ID, cmd)
// 			}
//
// 			if cmd == STOP {
// 				e.Running = false
// 				log.Printf("echo id: %v cmd: STOP aka %v\n", e.ID, cmd)
// 				goto loop
// 			}
//
// 			for {
// 				select {
// 				case cmd := <-e.Cmd:
// 					// log.Printf("echo: %v cmd: %v\n", e.ID, cmd)
// 					if cmd == EXIT {
// 						e.Initialised = false
// 						e.Running = false
// 						log.Printf("echo id: %v cmd: EXIT aka %v\n", e.ID, cmd)
// 						return
// 					}
//
// 					if cmd == STOP {
// 						e.Running = false
// 						log.Printf("echo id: %v cmd: STOP aka %v\n", e.ID, cmd)
// 						goto loop
// 					}
//
// 				case msg := <-e.In:
// 					// Print the message
// 					log.Println("echo id:", e.ID, "chan:", e.In, "msg:", msg)
// 				}
// 			}
// 		}
// 	}()
// }
//

// Stop ...
func (r *Router) Stop() {
	r.Lock()
	r.Running = false
	r.Unlock()
}
