package router

import "fmt"

// Router ...
type Router struct {
	ID     int //address
	Name   string
	Desc   string
	In     chan interface{}
	OutMap map[int]chan interface{} // out id = chan
}

// NewRouter ...
func NewRouter(id int, name, desc string) *Router {
	return &Router{
		id, //id
		// fmt.Sprintf("%v %05v", kind, id), //name
		name,
		desc,                           //desc
		make(chan interface{}),         //in
		make(map[int]chan interface{}), //out
	}
}

// PrintRouter ...
func PrintRouter(r *Router) {
	fmt.Println("id  :", r.ID)
	fmt.Println("name:", r.Name)
	fmt.Println("desc:", r.Desc)
	fmt.Println("in  :", r.In)
	for k, v := range r.OutMap {
		fmt.Println("out :", k, v)
	}
}
