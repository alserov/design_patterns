package main

import "fmt"

func main() {
	for num := range fib(1000) {
		fmt.Println(num)
	}
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
