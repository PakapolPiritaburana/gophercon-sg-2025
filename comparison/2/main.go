package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name string `json:"name"`
	Age  any    `json:"age"`
}

func main() {
	a := Person{"John", 30}

	aBs, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	var aUnmarchalled Person
	_ = json.Unmarshal(aBs, &aUnmarchalled)

	fmt.Printf("type of age %v\n", reflect.TypeOf(aUnmarchalled.Age))

	aEqual := reflect.DeepEqual(a, aUnmarchalled)

	if !aEqual {
		fmt.Printf("Not equal")
	} else {
		fmt.Printf("Equal\na is %v\naUnmarchalled is %v",
			a, aUnmarchalled)
	}
}
