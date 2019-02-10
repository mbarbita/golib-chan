package controller

type Controller struct {
	ID      int
	LoopMap map[int]*Loop
}

func NewController(id int) *Controller {
	return &Controller{
		ID:      id,
		LoopMap: make(map[int]*Loop),
	}
}

func (c *Controller) AddLoop(id int, loop *Loop) {
	c.LoopMap[id] = loop
}
