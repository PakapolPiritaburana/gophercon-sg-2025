package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type person struct {
	name string
	age  any
}

func main() {
	a := &person{"John", 30.00}

	aBs, _ := json.Marshal(a)
	var aUnmarchalled person
	_ = json.Unmarshal(aBs, &a)
	aEqual := reflect.DeepEqual(a, aUnmarchalled)

	if !aEqual {
		fmt.Printf("Not equal")
	} else {
		fmt.Printf("Equal\nmyDoc is %v\nmyDocUnmarchalled is %v",
			a, aUnmarchalled)
	}
}
