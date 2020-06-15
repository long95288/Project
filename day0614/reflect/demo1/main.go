package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2 = ", n2)

	fmt.Printf("rVal = %v rVal type = %T \n", rVal, rVal)
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2 = ", num2)

}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType = ", rType)

	rVal := reflect.ValueOf(b)

	iV := rVal.Interface()
	fmt.Printf("iv= %v iv type = %T \n", iV, iV)

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu Name = %v \n", stu.Name)
	}
}

// Student struct
type Student struct {
	Name string
	Age  int
}

func main() {
	stu := Student{
		Name: "tom",
		Age:  100,
	}
	reflectTest02(stu)
}
