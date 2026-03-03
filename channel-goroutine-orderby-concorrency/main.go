package main

import "fmt"

type result struct {
	value int
}

func worker(id int, ch chan<- result) {
	ch <- result{value: id * 2}
}

func main() {
	channels := make([]chan result, 5)

	for i := range channels {
		channels[i] = make(chan result, 1)
		go worker(i, channels[i])

	}

	for _, ch := range channels {
		r := <-ch
		fmt.Println(r.value)
	}
}
