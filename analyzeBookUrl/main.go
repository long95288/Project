package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "strings"
)

func readList(url string) []string {
    data,err := ioutil.ReadFile(url)
    list := []string{}
    if err != nil {
        log.Fatal(err)
    }
    err = json.Unmarshal(data,&list)
    if err != nil{
        log.Fatal(err)
    }
    return list
}
func main() {
    pattern := "list1 (%d).json"
    allList := []string{}
    for i:=1;i<=13;i++{
        url := fmt.Sprintf(pattern,i)
        fmt.Println("url = ",url)
        allList = append(allList,readList(url)...)
    }
    urlMap := make(map[string]struct{})
    for _,url := range allList{
        urlMap[url] = struct{}{}
    }
    content := ""
    for k,_ := range urlMap {
        if strings.Contains(k,".pdf") {
            content = content + k + "\n"
        }
    }
    ioutil.WriteFile("list.txt",[]byte(content),0666)
}
