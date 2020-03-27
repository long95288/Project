package main

import "fmt"

func main() {
	nums := []int{2,3,4,5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:",sum)

	for i,num := range nums{
		fmt.Printf("index = %d value = %d \n",i,num)
	}

	kvs := map[string]string{"a":"apple","b":"banana"}
	for k, v:=range kvs {
		fmt.Printf("%s -> %s \n",k,v)
	}
	// map[key type] value type {initial data}
	kvs2 := map[int]string{1:"hello",2:"world"}

	for k,v:=range kvs2{
		fmt.Printf("key = %d value=%s \n",k,v)
	}


	for i,c := range "golang"{
		fmt.Println(i,c)
	}
}

