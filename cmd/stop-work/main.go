package main

import "fmt"

func main() {
	work := make(chan int, 50)
	for i := range cap(work) {
		work <- i
	}

	stop := make(chan struct{})
	close(stop)
	fmt.Println("stop signal sent")

	loop(work, stop)
}

func loop[T any](work chan T, stop chan struct{}) {
	defer fmt.Println("work loop stopped")

	for {
		select {
		case <-stop:
			return
		case w := <-work:
			fmt.Println("processing message", w)
		}
	}
}
