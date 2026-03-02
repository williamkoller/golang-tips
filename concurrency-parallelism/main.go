package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("[%s] passo %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go task("A")
	go task("B")

	time.Sleep(500 * time.Millisecond)
}
