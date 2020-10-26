package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    input,_ := in.ReadString('\n')
    target,_ := in.ReadString('\n')
    
    map1 := make(map[byte]int)
    for _,v := range input {
        map1[byte(v)] += 1
    }
    fmt.Println(string(input))
    v,_ := map1[target[0]]
    v2,_ := map1[target[0] + 32]
    v3,_ := map1[target[0] - 32]
    fmt.Println(v + v2 + v3)
}
