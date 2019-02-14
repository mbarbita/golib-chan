package main

import (
	"fmt"
)

type Controller struct {
	loopMap map[int]Loop
}

type Loop struct {
	compMap map[int]Component
	cmdMap  map[int]chan int8
	inMap   map[int]chan interface{}
	outMap  map[int]chan interface{}
}

type Component interface {
	setID(int)
}

type Frame struct {
	id          int
	initialized bool
	running     bool
}

func (f *Frame) setID(id int) {
	f.id = id
	fmt.Printf("<frame> id on frame %+v set\n", f)
}

type Router struct {
	Frame
	outMap map[int]int
}

type Echo struct {
	Frame
}

func main() {

	controller1 := new(Controller)
	fmt.Printf("controller1: %+v\n", controller1)

	controller1.loopMap = make(map[int]Loop)
	controller1.loopMap[0] = Loop{
		compMap: make(map[int]Component),
		cmdMap:  make(map[int]chan int8),
		inMap:   make(map[int]chan interface{}),
		outMap:  make(map[int]chan interface{}),
	}
	fmt.Printf("controller1.loopMap: %+v\n", controller1.loopMap)

	fmt.Printf("controller1.loopMap[0].inMap: %+v\n", controller1.loopMap[0].inMap)
	fmt.Printf("controller1.loopMap[0].cmdMap: %+v\n", controller1.loopMap[0].cmdMap)
	fmt.Printf("controller1.loopMap[0].cmdMap: %+v\n", controller1.loopMap[0].outMap)

	c1 := controller1.loopMap[0]

	fmt.Printf("c1: %+v\n", c1)

	c1.compMap[0] = new(Router)
	c1.compMap[1] = new(Echo)
	fmt.Printf("c1: %+v\n", c1)
	fmt.Printf("c1.compMap[0]: %+v\n", c1.compMap[0])

	for k, v := range c1.compMap {
		v.setID(k)
	}
	fmt.Printf("c1: %+v\n", c1)
	fmt.Printf("c1.compMap[0]: %+v\n", c1.compMap[0])
}
