package main

import "fmt"

func main() {
    num1 := 100
    fmt.Printf("Type %T Value= %v \n",num1,num1)
    
    num2 := new(int)
    *num2 = 100
    fmt.Printf("Type %T value = %v *value = %v",num2,num2,*num2)
    
}
