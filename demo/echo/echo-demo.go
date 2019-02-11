package main

import (
	"fmt"
	"log"
	"time"

	ccore "github.com/mbarbita/golib-controller/core"
)

func main() {
	wch := make(chan bool)

	e1 := ccore.NewEcho(0)
	e2 := ccore.NewEcho(1)

	r1 := ccore.NewRouter(0, make(chan int8), make(chan interface{}))
	r1.ModOut(0, e1.In)
	r1.ModOut(1, e2.In)

	// e1 := ccore.NewEcho(0, make(chan int8), r1.OutMap[0])
	// e2 := ccore.NewEcho(1, make(chan int8), r1.OutMap[1])
	ccore.PrintRouter(r1)
	fmt.Println()
	ccore.PrintComp(e1)
	fmt.Println()
	ccore.PrintComp(e2)
	fmt.Println()

	r1.Start()

	e1.Init()
	e1.Run()
	e2.Init()
	e2.Run()

	// time.Sleep(5 * time.Second)

	go func() {
		time.Sleep(2000 * time.Millisecond)
		e1.Cmd <- ccore.STOP
	}()

	go func() {
		for i := 0; i < 5; i++ {
			log.Println("sending data:")
			r1.In <- 123
			time.Sleep(1000 * time.Millisecond)

			r1.In <- "blabla"
			time.Sleep(1000 * time.Millisecond)
		}
		time.Sleep(1000 * time.Millisecond)

		ccore.PrintRouter(r1)
		ccore.PrintComp(e1)
		ccore.PrintComp(e2)

		wch <- true

		// e1.Stop()
		// r1.Stop()
	}()

	<-wch

}
