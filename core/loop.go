package controller

type Loop struct {
	// ID      int
	CompMap map[int]interface{}
	ChanMap map[int]chan interface{}
}

func NewLoop() *Loop {
	return &Loop{
		// ID:      id,
		CompMap: make(map[int]interface{}),
		ChanMap: make(map[int]chan interface{}),
	}
}

func (l Loop) AddComp(id int, comp interface{}) {
	l.CompMap[id] = comp
}

func (l Loop) AddChan(id int, ch chan interface{}) {
	l.ChanMap[id] = ch
}
