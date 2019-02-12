package main

import (
	"log"
	"time"

	ccore "github.com/mbarbita/golib-controller/core"
)

func main() {
	wch := make(chan bool)
	echoMap := make(map[int]*ccore.Echo)
	echoMap[1] = ccore.NewEcho(1)
	echoMap[2] = ccore.NewEcho(2)
	// e1 := ccore.NewEcho(0)
	// e2 := ccore.NewEcho(1)

	r1 := ccore.NewRouter(0)
	for k, v := range echoMap {
		r1.ModOut(k, v.In)
	}
	// r1.ModOut(0, e1.In)
	// r1.ModOut(1, e2.In)

	ccore.PrintRouter(r1)
	for k, _ := range echoMap {
		ccore.PrintEcho(echoMap[k])
		// ccore.PrintEcho(echoMap[1])
	}

	r1.Init()
	r1.Run()

	for _, v := range echoMap {
		// r1.ModOut(k, v.In)
		v.Init()
		v.Run()
	}
	// e1.Init()
	// e1.Run()
	// e2.Init()
	// e2.Run()

	// time.Sleep(5 * time.Second)

	go func() {
		time.Sleep(2000 * time.Millisecond)
		echoMap[1].Stop()
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
		for k, _ := range echoMap {
			ccore.PrintEcho(echoMap[k])
			// ccore.PrintEcho(echoMap[1])
		}

		// ccore.PrintRouter(r1)
		// ccore.PrintEcho(echoMap[0])
		// ccore.PrintEcho(echoMap[1])

		// ccore.PrintRouter(r1)
		// ccore.PrintEcho(e1)
		// ccore.PrintEcho(e2)

		wch <- true

		// e1.Stop()
		// r1.Stop()
	}()

	<-wch

}
