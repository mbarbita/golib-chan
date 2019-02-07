package comp

import "fmt"

type Echo struct {
	ID      int
	In      chan interface{}
	Running bool
}

func NewEcho(id int, ch chan interface{}) *Echo {
	return &Echo{
		id,
		ch,
		false,
	}
}

func PrintComp(e *Echo) {
	fmt.Println("Echo:")
	fmt.Println("id     :", e.ID)
	fmt.Println("running:", e.Running)
	fmt.Println("in chan:", e.In)
	fmt.Println()
}

func (e *Echo) Start() {
	go func() {
		e.Running = true
		for {
			if !e.Running {
				// return
				break
			}
			msg := <-e.In
			// Print the message
			// fmt.Println(msg)
			fmt.Println("echo id:", e.ID, "chan:", e.In, "msg:", msg)
		}
	}()
}

// Stop ...
func (e *Echo) Stop() {

	e.Running = false

}
