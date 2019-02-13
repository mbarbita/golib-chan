package controller

import (
	"fmt"
	"log"
)

// Router ...
type Router struct {
	*Frame
	OutMap map[int]chan interface{} // out id = chan
}

func (r *Router) InMsg(inMsg interface{}) {
	// fmt.Println(inMsg)
	// log.Println("*** frame id:", r.ID, "chan:", r.In, "msg:", inMsg)

	for k, v := range r.OutMap {
		select {
		case v <- inMsg:
		default:
			log.Printf("router %v could not send to chan id: %v chan:%v\n",
				r.ID, k, v)
		}
	}
}

// NewRouter ...
func NewRouter(id int) *Router {
	r := &Router{
		Frame:  NewFrame(id),
		OutMap: make(map[int]chan interface{}), //out
	}
	// r.SetFn(r.InMsg)
	r.Fn = r.InMsg
	return r
}

// PrintRouter ...
func PrintRouter(r *Router) {
	fmt.Println("Router:")
	PrintFrame(r.Frame)
	for k, v := range r.OutMap {
		fmt.Printf("out id : %v, chan: %v type: %T\n", k, v, v)
	}
	fmt.Println()
}

// ModifyOut ...
func (r *Router) ModOut(outID int, ch chan interface{}) {
	r.Lock()
	r.OutMap[outID] = ch
	r.Unlock()
}
