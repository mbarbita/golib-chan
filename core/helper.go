package controller

import (
	"fmt"
	"sort"
	"time"
)

type Dur struct {
	In       chan time.Duration
	DurSlice []time.Duration
}

func NewDur() *Dur {
	return &Dur{
		In: make(chan time.Duration),
	}
}

func (d *Dur) Run() {
	go func() {
		// d.In = make(chan time.Duration)
		// DurSlice = make([]time.Duration)
		for {
			d.DurSlice = append(d.DurSlice, <-d.In)
		}
	}()
}

func (d *Dur) PrintDur() {
	// sort.Sort(d.DurSlice)

	sort.Slice(d.DurSlice, func(i, j int) bool {
		return d.DurSlice[i] < d.DurSlice[j]
	})

	for _, v := range d.DurSlice {
		fmt.Println(v)

	}
}
