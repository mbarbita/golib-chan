package router

import (
	"fmt"
	"sync"
)

// Router ...
type Router struct {
	ID      int //address
	Name    string
	Desc    string
	Running bool
	In      chan interface{}
	OutMap  map[int]chan interface{} // out id = chan
}

// NewRouter ...
func NewRouter(id int, name, desc string) *Router {
	return &Router{
		id, //id
		// fmt.Sprintf("%v %05v", kind, id), //name
		name,
		desc, //desc
		false,
		make(chan interface{}),         //in
		make(map[int]chan interface{}), //out
	}
}

// PrintRouter ...
func PrintRouter(r *Router) {
	fmt.Println("id     :", r.ID)
	fmt.Println("name   :", r.Name)
	fmt.Println("desc   :", r.Desc)
	fmt.Println("running:", r.Running)
	fmt.Println("in     :", r.In)
	for k, v := range r.OutMap {
		fmt.Println("out id:", k, "out chan:", v)
	}
}

// ModifyOut ...
func (r *Router) ModifyOut(cmd int, id int, ch chan interface{}) {

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

	switch cmd {
	case 1: // Add or modify destination
		r.OutMap[id] = ch
	// case 2:
	// 	r.OutMap[id] = ch
	case 2: // Delete destination
		delete(r.OutMap, id)
		close(ch)
	}
}

// ModifyIn ...
func (r *Router) ModifyIn(ch chan interface{}) {

	mutex := &sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()
	r.In = ch
	close(r.In)
}

// ChangeID ...
func (r *Router) ChangeID(id int) {

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
				return
			}
			msg := <-r.In
			// loop trough client map and send the message
			for _, v := range r.OutMap {
				v <- msg
			}
			// log.Printf("sent message to %d dst", len(bcr.dst))
			// }
		}
	}()
}

// Stop ...
func (r *Router) Stop() {

	r.Running = false

}
