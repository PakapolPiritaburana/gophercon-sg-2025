package main

import (
	"fmt"
	"reflect"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Employee struct {
	id      int
	Name    string
	Salary  float64
	Company string
}

func CreateEmployee(id int, name string, salary float64, company string) *Employee {
	return &Employee{id, name, salary, company}
}

func main() {
	got := CreateEmployee(1, "John", 1000.00, "Google")
	want := &Employee{2, "John", 1000.00, "Google"}

	ignore := cmpopts.IgnoreUnexported(Employee{})
	//ignore := cmpopts.IgnoreFields(Employee{}, "Id")
	fmt.Println(cmp.Equal(got, want, ignore))

	fmt.Printf("got and want are equal: %t\n",
		reflect.DeepEqual(got.id, want.id))
}
