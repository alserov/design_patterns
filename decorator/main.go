package main

import "fmt"

var (
	base  = PC{}
	extra = ExtraPC{
		Cpu:     8,
		Wrapper: base,
	}
	ultra = UltraPC{
		Cpu:     8,
		SSD:     16,
		Wrapper: base,
	}
)

func main() {
	fmt.Println(base.GetPrice())
	fmt.Println(extra.GetPrice())
	fmt.Println(ultra.GetPrice())
}

type Wrapper interface {
	GetPrice() float64
}

type PC struct{}

func (P PC) GetPrice() float64 {
	return 10
}

type ExtraPC struct {
	Cpu     int
	Wrapper Wrapper
}

func (e ExtraPC) GetPrice() float64 {
	return e.Wrapper.GetPrice() * float64(e.Cpu)
}

type UltraPC struct {
	Cpu     int
	SSD     int
	Wrapper Wrapper
}

func (u UltraPC) GetPrice() float64 {
	return u.Wrapper.GetPrice() * float64(u.Cpu) * float64(u.SSD)
}
