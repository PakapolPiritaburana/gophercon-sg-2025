package main

import (
	"fmt"
	"reflect"
)

func main() {
	sliceA := []int{1, 2}
	sliceB := []int{1, 2}

	//if sliceA == sliceB {
	//	fmt.Println("== Equal")
	//} else {
	//	fmt.Println("== Not equal")
	//}

	if reflect.DeepEqual(sliceA, sliceB) {
		fmt.Println("DeepEqual Equal")
	} else {
		fmt.Println("DeepEqual Not equal")
	}
}
