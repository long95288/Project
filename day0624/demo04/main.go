package main

import (
    "fmt"
    "html/template"
    "os"
)

func main() {
    tEmpty := template.New("template test")
    tEmpty = template.Must(tEmpty.Parse("空  if demo:{{if ``}} 不会输出的内容,{{end}}"))
    tEmpty.Execute(os.Stdout,nil)
    
    fmt.Println()
    tWithValue := template.New("template test")
    tWithValue = template.Must(tWithValue.Parse("不为空 if demo:{{if `value`}} 输出内容{{end}}"))
    tWithValue.Execute(os.Stdout,nil)
    
    tIfElse := template.New("template test")
    tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `ff`}} if true context {{else}} else 内容 {{end}}"))
    tIfElse.Execute(os.Stdout,nil)
}
