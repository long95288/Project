package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method :" ,r.Method)
    if r.Method == "GET"{
        t,_ := template.ParseFiles("login.html")
        
        t.Execute(w,nil)
    }else{
        r.ParseForm()
        fmt.Println("username",r.Form["username"])
        fmt.Println("password",r.Form["password"])
    }
}
func main() {
    http.HandleFunc("/login",login)
    err := http.ListenAndServe(":9090",nil)
    if err != nil {
        log.Fatal("listen err ",err)
    }
}
