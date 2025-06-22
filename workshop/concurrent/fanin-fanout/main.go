package main

import (
	"fmt"
)

func main() {
	r := fanOutFanIn([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	fmt.Println(r)

}

// Go example
func fanOutFanIn(input []int) int {
	// Fan Out - กระจายงานไปให้ goroutines
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		sum := 0
		for i := 0; i < len(input)/2; i++ {
			sum += input[i]
		}
		ch1 <- sum
	}()

	go func() {
		sum := 0
		for i := len(input) / 2; i < len(input); i++ {
			sum += input[i]
		}
		ch2 <- sum
	}()

	// Fan In - รวบรวมผลลัพธ์
	return <-ch1 + <-ch2
}
