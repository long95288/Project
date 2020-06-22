package main

import (
    "fmt"
    "log"
    "net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    for k,v := range r.Form {
        fmt.Println("Key: ",k)
        fmt.Println("Value: ",v)
    }
    fmt.Fprint(w,"Hello astaxie")
}
func main() {
    http.HandleFunc("/",sayHello)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("Listen err",err)
    }
    
    
}
