package controller

import (
	"fmt"
	"log"
	"sync"
)

// Router ...
type Router struct {
	ID      int //address
	Running bool
	sync.Mutex
	In     chan interface{}
	OutMap map[int]chan interface{} // out id = chan
}

// NewRouter ...
func NewRouter(id int, inCh chan interface{}) *Router {
	return &Router{
		ID:      id,                             //id
		Running: false,                          // running
		In:      inCh,                           //in
		OutMap:  make(map[int]chan interface{}), //out
	}
}

// PrintRouter ...
func PrintRouter(r *Router) {
	fmt.Println("Router:")
	// fmt.Printf("%-10v:\n", "Router")
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

// Stop ...
func (r *Router) Stop() {
	r.Lock()
	r.Running = false
	r.Unlock()
}
