package main

import (
	"log"
	"time"

	ccore "github.com/mbarbita/golib-controller/core"
)

func main() {
	// frameID := 0
	wch := make(chan bool)
	echoMap := make(map[int]*ccore.Echo)

	r1 := ccore.NewRouter(0)
	for frameID := 1; frameID <= 1000; frameID++ {
		echoMap[frameID] = ccore.NewEcho(frameID)
		// frameID++
		// echoMap[frameID] = ccore.NewEcho(frameID)
		// frameID++
	}

	// frameID++
	for k, v := range echoMap {
		r1.ModOut(k, v.In)
	}

	// print stopped state
	if false {
		ccore.PrintRouter(r1)
		for k := range echoMap {
			ccore.PrintEcho(echoMap[k])
		}
	}

	r1.Init()
	r1.Run()

	for _, v := range echoMap {

		v.Init()
		v.Run()
	}

	time.Sleep(time.Second)

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

		if false {
			ccore.PrintRouter(r1)
			for k := range echoMap {
				ccore.PrintEcho(echoMap[k])
			}
		}

		wch <- true

		// e1.Stop()
		// r1.Stop()
	}()

	<-wch

}
