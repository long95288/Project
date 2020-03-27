package main

import "fmt"

/**
go map的增删改查
*/
func main() {
	countryCapitalMap := map[string]string{"France":"Paris","Italy":"Rome"}
	fmt.Println("初始化的map")
	printMap(countryCapitalMap)
	// 添加
	countryCapitalMap["Japan"] = "Tokyo"
	fmt.Println("添加数据")
	printMap(countryCapitalMap)
	// 删除数据
	fmt.Println("删除数据")
	delete(countryCapitalMap,"France")
	printMap(countryCapitalMap)
	// 删除不存在的数据
	fmt.Println("删除不存在的数据")
	delete(countryCapitalMap,"American")
	printMap(countryCapitalMap)
	// 更改数据
	fmt.Println("更改数据")
	countryCapitalMap["Japan"] = "tokyo"
	printMap(countryCapitalMap)
	fmt.Println("更改不存在的数据,会直接添加新的数据")
	// 华盛顿
	countryCapitalMap["American"] = "Washington"
	printMap(countryCapitalMap)
}
func printMap(m map[string]string){
	for item := range m{
		fmt.Println(item,"首都是",m[item])
	}
}
