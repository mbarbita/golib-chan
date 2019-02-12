package controller

import (
	"fmt"
	"log"
)

type Echo struct {
	*Frame
}

func (e *Echo) InMsg(inMsg interface{}) {
	log.Println("*** echo id:", e.ID, "chan:", e.In, "msg:", inMsg)
}

func NewEcho(id int) *Echo {
	e := &Echo{
		Frame: NewFrame(id),
	}
	e.SetFn(e.InMsg)
	return e
}

func PrintEcho(e *Echo) {
	fmt.Println("Echo:")
	PrintFrame(e.Frame)
	fmt.Println()
}
