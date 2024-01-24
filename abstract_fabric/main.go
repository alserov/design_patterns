package main

import (
	"errors"
	"fmt"
)

func main() {
	m1, _ := NewFactory(M1)
	m1.GetDisplay().PrintDetails()
	m1.GetComputer().PrintDetails()

	m2, _ := NewFactory(M2)
	m2.GetDisplay().PrintDetails()
	m2.GetComputer().PrintDetails()

	_, err := NewFactory("m3")
	fmt.Println(err)
}

const (
	M1 = "m1"
	M2 = "m2"
)

type Computer interface {
	PrintDetails()
}

type Display interface {
	PrintDetails()
}

// Factory
type Factory interface {
	GetComputer() Computer
	GetDisplay() Display
}

func NewFactory(t string) (Factory, error) {
	switch t {
	case M1:
		return &FactoryM1{}, nil
	case M2:
		return &FactoryM2{}, nil
	default:
		return nil, errors.New("invalid type")
	}
}

// Factory M1
type FactoryM1 struct {
}

func (f FactoryM1) GetComputer() Computer {
	return &ComputerM1{
		Cpu:    8,
		Memory: 16,
	}
}

func (f FactoryM1) GetDisplay() Display {
	return &DisplayM1{
		Size: 24,
	}
}

// Factory M2
type FactoryM2 struct {
}

func (f FactoryM2) GetComputer() Computer {
	return &ComputerM2{
		Cpu:    4,
		Memory: 8,
	}
}

func (f FactoryM2) GetDisplay() Display {
	return &DisplayM2{
		Size: 32,
	}
}

// M1
type ComputerM1 struct {
	Memory int
	Cpu    int
}

func (c ComputerM1) PrintDetails() {
	fmt.Printf("Mem: %d \t Cpu: %d \n", c.Memory, c.Cpu)
}

type DisplayM1 struct {
	Size int
}

func (d DisplayM1) PrintDetails() {
	fmt.Printf("Size: %d \n", d.Size)
}

// M2
type ComputerM2 struct {
	Memory int
	Cpu    int
}

func (c ComputerM2) PrintDetails() {
	fmt.Printf("Mem: %d \t Cpu: %d \n", c.Memory, c.Cpu)
}

type DisplayM2 struct {
	Size int
}

func (d DisplayM2) PrintDetails() {
	fmt.Printf("Size: %d \n", d.Size)
}
