package main

import "fmt"

func main() {
    a := 1
    a ++
    fmt.Printf("%v \n",a)
    
    p := &a
    *p ++
    fmt.Printf("%v \n",a)
    
}
