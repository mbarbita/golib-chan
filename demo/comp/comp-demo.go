package main

import (
	"fmt"
	"time"

	"github.com/mbarbita/golib-chan/comp"
	"github.com/mbarbita/golib-chan/router"
)

func main() {

	r1 := router.NewRouter(0, make(chan interface{}))
	r1.ModOut(0, make(chan interface{}, 1))
	r1.ModOut(1, make(chan interface{}, 1))

	e1 := comp.NewEcho(0, r1.OutMap[0])
	e2 := comp.NewEcho(1, r1.OutMap[1])
	router.PrintRouter(r1)
	comp.PrintComp(e1)
	comp.PrintComp(e2)

	r1.Start()

	e1.Start()
	e2.Start()

	// time.Sleep(5 * time.Second)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending data:")
			r1.In <- 123
			time.Sleep(time.Second)

			r1.In <- "blabla"
			time.Sleep(time.Second)
		}
		// e1.Stop()
		// r1.Stop()
	}()

	// select {}
	// wch := make(chan bool)
	// <-wch
	time.Sleep(10 * time.Second)

}
