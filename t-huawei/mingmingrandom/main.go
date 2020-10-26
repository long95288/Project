package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    in := bufio.NewReader(os.Stdin)
    inputArray := []int{}
    for str,err := in.ReadString('\n'); err == nil;str,err = in.ReadString('\n'){
        str = strings.Replace(str,"\r\n","", -1)
        str = strings.Replace(str, "\n", "", -1)
        if "" == str{
            break
        }
        //fmt.Printf("out = %q.\n", str)
        i,_ := strconv.Atoi(str)
        inputArray = append(inputArray, i)
    }
    // 3 1 2 3 2 1 2
    for i := 0; i < len(inputArray);{
        tmp := inputArray[i + 1 : i + inputArray[i] + 1]
        sort.Ints(tmp)
        index := 0
        fmt.Println(tmp[index])
        for j := 1;j<len(tmp);j++{
            if tmp[index] != tmp[j] {
                index ++
                tmp[index] = tmp[j]
                fmt.Println(tmp[index])
            }
        }
        i = i + inputArray[i] + 1
    }
}
