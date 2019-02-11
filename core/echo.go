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
	return &Echo{
		Frame: NewFrame(id),
	}
}

func PrintComp(e *Echo) {
	fmt.Println("Echo:")
	PrintFrame(&e.Frame)
	// fmt.Println()
}
