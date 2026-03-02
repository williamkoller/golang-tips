package main

import "time"

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				println("Tick at", t.String())
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	println("Ticker stopped")
}
