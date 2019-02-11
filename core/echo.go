package controller

import (
	"fmt"
)

type Echo struct {
	Frame
}

func (e *Echo) InMsg(inMsg interface{}) {
	fmt.Println(inMsg)
}

func NewEcho(id int) *Echo {

	e := &Echo{
		Frame: NewFrame(id),
	}
	e.AddFn(0, e.InMsg)
	return e
}

func PrintComp(e *Echo) {
	fmt.Println("Echo:")
	PrintFrame(&e.Frame)
	// fmt.Println()
}
