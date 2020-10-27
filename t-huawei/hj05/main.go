package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func parseStr(s string)  {
    var hexMap = map[rune]int{
        '1':1,
        '2':2,
        '3':3,
        '4':4,
        '5':5,
        '6':6,
        '7':7,
        '8':8,
        '9':9,
        '0':0,
        'a': 10,
        'A': 10,
        'b': 11,
        'B': 11,
        'C': 12,
        'c': 12,
        'd': 13,
        'D': 13,
        'e': 14,
        'E': 14,
        'f': 15,
        'F': 15,
    }
    base := len(s)
    ret := 0
    for _, v := range s {
        tmp := 1
        for i := 1; i < base; i++ {
            tmp *= 16
        }
        d,_ := hexMap[v]
        ret += d * tmp
        base --
    }
    fmt.Println(ret)
}

func main() {
   in := bufio.NewReader(os.Stdin)
   for str,err := in.ReadString('\n');err == nil;str,err = in.ReadString('\n'){
        str = strings.Replace(str, "\r\n", "", -1)
        str = strings.Replace(str, "\n","", -1)
        if str == "" {
            break
        }
        parseStr(string(str[2:]))
   }
}
