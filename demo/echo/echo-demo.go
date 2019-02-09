package main

import (
	"fmt"
	"time"

	ccore "github.com/mbarbita/golib-controller/core"
)

func main() {
	wch := make(chan bool)

	r1 := ccore.NewRouter(0, make(chan interface{}))
	r1.ModOut(0, make(chan interface{}, 1))
	r1.ModOut(1, make(chan interface{}, 1))

	e1 := ccore.NewEcho(0, r1.OutMap[0])
	e2 := ccore.NewEcho(1, r1.OutMap[1])
	ccore.PrintRouter(r1)
	ccore.PrintComp(e1)
	ccore.PrintComp(e2)

	r1.Start()

	e1.Start()
	e2.Start()

	// time.Sleep(5 * time.Second)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("sending data:")
			r1.In <- 123
			// time.Sleep(100 * time.Millisecond)

			r1.In <- "blabla"
			// time.Sleep(100 * time.Millisecond)
		}
		time.Sleep(1000 * time.Millisecond)
		wch <- true

		// e1.Stop()
		// r1.Stop()
	}()

	<-wch

}
