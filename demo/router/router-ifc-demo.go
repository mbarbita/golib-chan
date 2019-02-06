package main

import (
	"fmt"

	"github.com/mbarbita/golib-chan/router"
)

func main() {

	r2 := router.NewRouter(0, "some router", "kitchen")
	r2.OutMap[0] = make(chan interface{})
	fmt.Println(r2)

	go func() {
		fmt.Println("sending data:")
		r2.OutMap[0] <- 123
		r2.OutMap[0] <- "blabla"
	}()

	router.PrintRouter(r2)
	ifc := <-r2.OutMap[0]
	fmt.Printf("reading from chan: val: %v, type: %T\n", ifc, ifc)

	ifc = <-r2.OutMap[0]
	fmt.Printf("reading from chan: val: %v, type: %T\n", ifc, ifc)
}
