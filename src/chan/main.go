package main

import (
	"sync"
	"fmt"
)

type ThreadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (m *ThreadSafeSet) Inter() <-chan interface{} {
	ch := make(chan interface{}, len(m.s))
	go func() {
		m.RLock()

		for elem, value := range m.s {
			ch <- value
			fmt.Println("Iter:", elem, value)
		}

		close(ch)
		m.RUnlock()
	}()

	return ch
}

func main()  {
	ts := ThreadSafeSet{
		s: []interface{} {
			"1", "2",
		},
	}
	v := ts.Inter()
	for {
		i, ok := <-v
		if !ok {
			break
		}
		fmt.Println(i)
	}
}