package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age`
	Score float32 `json:"成绩"`
	Sex   string
}

func (s Monster) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}
func (s Monster) Print() {
	fmt.Println("====start====")
	fmt.Println(s)
	fmt.Println("====end====")
}
func testStruct(a interface{}) {

	rType := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	kd := rVal.Kind()
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	num := rVal.NumField()
	fmt.Printf("struct has %d fields \n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v", i, rVal.Field(i))
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d tag = %v \n", i, tagVal)
		}
	}

	numOfMethod := rVal.NumMethod()
	fmt.Printf("strcut has %d methods \n", numOfMethod)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := rVal.Method(0).Call(params)
	fmt.Println("res = ", res[0].Int())

}
func main() {
	var a Monster = Monster{
		"飞鼠", 300, 23.0, "雌性",
	}
	testStruct(a)

}
