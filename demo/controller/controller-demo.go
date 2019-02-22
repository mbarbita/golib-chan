package main

import (
	"fmt"

	ccore "github.com/mbarbita/golib-controller/core"
)

func main() {
	c1 := ccore.NewController(0)

	c1.AddLoop(0, ccore.NewLoop())
	c1.AddLoop(1, ccore.NewLoop())

	c1.LoopMap[0].AddComp(0, ccore.NewRouter(0))

	fmt.Println(c1)
	// fmt.Println(c1.LoopMap[0].CompMap[0])
	fmt.Println("controler", c1.ID, "LoopMap:")
	for k, v := range c1.LoopMap {
		fmt.Println("key:", k, "val:", v)
	}
}
