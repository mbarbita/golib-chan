package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type Loop struct {
	compMap map[int]Component
}

type Component interface {
	run(chan interface{}, chan int8)
	setID(int)
}

type Frame struct {
	id          int
	initialized bool
	running     bool
	cmd         chan int8
}

func (f *Frame) setID(id int) {
	fmt.Printf("frame: %+v, set id: %v\n", f, id)
	f.id = id
}

func (f *Frame) run(in chan interface{}, cmd chan int8) {
	fmt.Printf("frame: %+v - run\n", f)
}

type Router struct {
	Frame
	outMap map[int]interface{}
}

type Echo struct {
	Frame
}

func (e *Echo) run(in chan interface{}, cmd chan int8) {
	e.cmd = cmd
	e.initialized = true
	e.running = true
	// fmt.Printf("echo id: %+v - run\n", e.id)
	go func() {
		for {
			fmt.Println("Waiting:", in, "id:", e.id)
			msg := <-in
			fmt.Printf("received %v:\n", msg)
		}
	}()
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func main() {

	inMap := make(map[int]chan interface{})

	loop := new(Loop)
	loop.compMap = make(map[int]Component)

	loop.compMap[0] = new(Router)
	loop.compMap[1] = new(Echo)

	fmt.Println()

	for k, v := range loop.compMap {
		v.setID(k)
		inMap[k] = make(chan interface{})
		v.run(inMap[k], make(chan int8))
		fmt.Printf("comp id %v, val: %+v\n", k, v)
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("in map:")
	fmt.Println(inMap)
	fmt.Println("comp map:")
	for k, v := range loop.compMap {
		// fmt.Println(loop.loopMap)
		fmt.Printf("k: %v, v: %+v\n", k, v)
	}
	fmt.Println()

	for i := 0; i < 5; i++ {
		inMap[1] <- "daaaaaaaaaa"
		time.Sleep(time.Second)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Gob Example")
	// student := Student{"Ketan Parmar", 35}
	err := writeGob("./loop.gob", loop)
	if err != nil {
		fmt.Println("write error:", err)
	}

	var loop2 = new(Loop)
	err = readGob("./loop.gob", loop2)
	if err != nil {
		fmt.Println("read error:", err)
	} else {
		fmt.Println(loop2)
	}

}
