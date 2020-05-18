package main

import (
    "fmt"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("now = %v now Type=%T",now,now)
    
    fmt.Println("year = ",now.Year())
    fmt.Println("mon = ",now.Month())
    fmt.Println("mon = ",int(now.Month()))
    fmt.Println("day = ",now.Day())
    fmt.Println("hour = ",now.Hour())
    fmt.Println("min = ",now.Minute())
    fmt.Println("second = ",now.Second())
    
    
    
}
