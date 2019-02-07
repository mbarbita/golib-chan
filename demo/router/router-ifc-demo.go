package main

import (
	"fmt"

	"github.com/mbarbita/golib-chan/router"
)

func main() {
	// inCh := make(chan interface{})
	// r1 := router.NewRouter(0, inCh)
	r1 := router.NewRouter(0, make(chan interface{}))
	// r1.OutMap[0] = make(chan interface{})
	r1.ModOut(0, make(chan interface{}))
	// r1.OutMap[1] = make(chan interface{})
	r1.ModOut(1, make(chan interface{}))
	router.PrintRouter(r1)

	go func() {
		fmt.Println("sending data:")
		r1.OutMap[0] <- 123
		r1.OutMap[0] <- "blabla"
	}()

	ifc := <-r1.OutMap[0]
	fmt.Printf("reading from chan: val: %v, type: %T\n", ifc, ifc)

	ifc = <-r1.OutMap[0]
	fmt.Printf("reading from chan: val: %v, type: %T\n", ifc, ifc)
}
