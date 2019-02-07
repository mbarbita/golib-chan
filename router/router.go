package router

import (
	"fmt"
	"log"
	"sync"
)

// Router ...
type Router struct {
	ID      int //address
	Running bool
	In      chan interface{}
	OutMap  map[int]chan interface{} // out id = chan
}

// NewRouter ...
func NewRouter(id int, inCh chan interface{}) *Router {
	return &Router{
		id,                             //id
		false,                          // running
		inCh,                           //in
		make(map[int]chan interface{}), //out
	}
}

// PrintRouter ...
func PrintRouter(r *Router) {
	fmt.Println("Router:")
	fmt.Println("id     :", r.ID)
	fmt.Println("running:", r.Running)
	fmt.Println("in chan:", r.In)
	for k, v := range r.OutMap {
		fmt.Printf("out id : %v, chan: %v\n", k, v)
	}
	fmt.Println()
}

// ModifyOut ...
func (r *Router) ModOut(outID int, ch chan interface{}) {

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	r.OutMap[outID] = ch
}

func (r *Router) DelOut(outID int, ch chan interface{}) {

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	delete(r.OutMap, outID)
	close(ch)
}

// ModifyIn ...
func (r *Router) ModIn(ch chan interface{}) {

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	r.In = ch
	close(r.In)
}

// ChangeID ...
func (r *Router) ModID(id int) {

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	r.ID = id

}

// Start ...
func (r *Router) Start() {
	go func() {
		r.Running = true
		for {
			if !r.Running {
				// return
				break
			}

			msg := <-r.In
			// loop trough client map and send the message
			for _, v := range r.OutMap {
				// v <- msg

				select {
				case v <- msg:
				// sent msg down chan and didn't block
				// case <-time.After(3 * time.Second):
				// 	log.Printf("Send timeout to chan: %v\n", v)
				default:
					// sent nothing and would have blocked
					log.Printf("could not send to chan: %v\n", v)
				}

			}
		}
		log.Printf("break: %v\n", r.ID)

	}()
}

// Stop ...
func (r *Router) Stop() {

	r.Running = false

}
