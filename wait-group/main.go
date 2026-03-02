package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("workder %d finalizado\n", id)
		}(i)
	}

	wg.Wait()

	fmt.Println("todas concluidas")
}
