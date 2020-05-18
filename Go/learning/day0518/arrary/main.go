package main

import "fmt"

func main() {
    var intArr [3]int
    
    fmt.Println(intArr)
    
    intArr[0] = 10
    intArr[1] = 20
    intArr[2] = 30
    fmt.Println(intArr)
    fmt.Printf("address = %p intArr[0] address = %p intArr[1] address = %p\n",
        &intArr,&intArr[0],&intArr[1])
    fmt.Printf("%T",intArr[2])
    var iptr *int
    iptr = &intArr[1]
    fmt.Printf("address = %p value = %v",iptr,*iptr)
    
    var numArr01 [3]int = [3]int{1,2,3}
    fmt.Println(numArr01)
    var numArr02 = [3]int{4,5,6}
    fmt.Println(numArr02)
    var numArr03 = [...]int{7,8,9}
    fmt.Println(numArr03)
    var numArr04 = [...]int{0:1,1:10,2:11,3:12}
    fmt.Println(numArr04)
    var numArr05 = [...]string{1:"Hello",2:"23214",0:"world"}
    fmt.Println(numArr05)
    var ptr *[3]int
    ptr = &numArr01
    fmt.Println(ptr,*ptr)
    num1 := ptr[0]
    fmt.Println(num1)
    
}
