package main

import (
	"fmt"
	"reflect"
)

func main() {
	//sliceA := []string{"a", "b", "c"}
	sliceA := []int{1, 2}
	//sliceA := make([]string, 3)
	//sliceA[0] = "a"
	//sliceA[1] = "b"
	//sliceA[2] = "c"
	sliceB := []int{1, 2}
	//sliceB := []string{"a", "b", "c"}

	if sliceA[0] == sliceB[0] {
		fmt.Println("Index 0 = Equal")
	} else {
		fmt.Println("Index 0 = Not equal")
	}

	if reflect.DeepEqual(sliceA, sliceB) {
		fmt.Println("DeepEqual = Equal")
	} else {
		fmt.Println(" DeepEqual = Not equal")
	}
}
