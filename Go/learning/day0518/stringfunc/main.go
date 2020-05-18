package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    str := "Hello北"
    fmt.Println("str lens = ",len(str))
    
    str2 := "English英文"
    for i,v := range []rune(str2){
        fmt.Printf("index = %d value = %c\n",i,v)
    }
    
    n,err := strconv.Atoi("200")
    if nil != err{
        fmt.Println("转换错误",err)
    }else{
        fmt.Println("转换的整数 ",n)
    }
    n,err = strconv.Atoi("Hello")
    if nil != err {
        fmt.Println("转换错误",err)
    }else{
        fmt.Println("转换成功",n)
    }
    
    str3 := strconv.Itoa(12455)
    fmt.Printf("str = %v,str = %T",str3,str3)
    
    str = string([]byte{97,98,99})
    fmt.Printf("str = %v\n",str)
    str = strconv.FormatInt(123,2)
    fmt.Printf("123 二进制 %v\n",str)
    str = strconv.FormatInt(123,16)
    fmt.Printf("123 16进制 %v\n",str)
    
    b := strings.Contains("seafood","food")
    fmt.Printf("b = %v\n",b)
    num := strings.Count("cecheee","e")
    
    fmt.Printf("num = %d\n",num)
    b = strings.EqualFold("abc","ABc")
    fmt.Printf("b = %v\n",b)
    fmt.Printf("%v\n","abc" == "Abc")
    
    index := strings.Index("NLT_abcabcabc","abc")
    fmt.Printf("index = %v \n",index)
    index = strings.LastIndex("go golang","go")
    fmt.Printf("index = %v \n",index)
    str2 = "go go hello"
    str = strings.Replace(str2,"go","Hello",-1)
    fmt.Println("str = ",str)
    
    strArr := strings.Split("Hell,World",",")
    for i,v := range strArr{
        fmt.Println("index = ",i,"value = ",v)
    }
    
    str = "goLang Hello"
    str = strings.ToLower(str)
    fmt.Printf("%v\n",str)
    str = strings.ToUpper(str)
    fmt.Printf("%v\n",str)
    
    str = strings.TrimSpace(" tn a lone gopher ntrn  ")
    fmt.Printf("str = %q\n",str)
    
    str = strings.Trim("!hello!","!")
    fmt.Println(str)
    str = strings.TrimLeft("! hello!"," !")
    fmt.Println(str)
    str = strings.TrimRight("! hello!"," !")
    fmt.Println(str)
    fmt.Println(strings.HasPrefix("ftp://192.168.10.1","ftp"))
    fmt.Println(strings.HasSuffix("fagaa.jpg","jpg"))
}
