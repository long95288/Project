package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    input,_ := in.ReadString('\n')
    arr := strings.Split(input, " ")
    fmt.Println(len(strings.Replace(arr[len(arr) - 1], "\n", "", -1)))
}
