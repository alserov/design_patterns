package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	s := eventSubject{
		observers: sync.Map{},
	}

	var (
		o1 = eventObserver{id: 1, time: time.Now()}
		o2 = eventObserver{id: 2, time: time.Now()}
	)

	s.AddListener(&o1)
	s.AddListener(&o2)

	go func() {
		select {
		case <-time.After(time.Millisecond):
			s.RemoveListener(&o1)
		}
	}()

	for x := range fib(1000) {
		s.Notify(Event{data: x})
	}
}

type (
	Event struct {
		data int
	}

	Observer interface {
		NotifyCallback(event Event)
	}

	Subject interface {
		AddListener(observer Observer)
		RemoveListener(observer Observer)
		Notify(event Event)
	}

	eventObserver struct {
		id   int
		time time.Time
	}

	eventSubject struct {
		observers sync.Map
	}
)

func (e *eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Received: %v ID: %d time: %v \n ", event.data, e.id, time.Since(e.time))
}

func (e *eventSubject) AddListener(obs Observer) {
	e.observers.Store(obs, struct{}{})
}

func (e *eventSubject) RemoveListener(obs Observer) {
	e.observers.Delete(obs)
}

func (e *eventSubject) Notify(eve Event) {
	e.observers.Range(func(key, value any) bool {
		if key == nil || value == nil {
			return false
		}

		key.(Observer).NotifyCallback(eve)
		return true
	})
}

func fib(n int) chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()

	return out
}
