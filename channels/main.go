package main

import "fmt"

func sum(nums []int, ch chan int) {
	total := 0

	for _, n := range nums {
		total += n
	}

	ch <- total // envia o resultado
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan int)

	go sum(nums[:3], ch)
	go sum(nums[3:], ch)

	a, b := <-ch, <-ch

	fmt.Println(a + b)
}
