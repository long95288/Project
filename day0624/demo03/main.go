package main

import (
    "html/template"
    "os"
)

type Friend struct {
    FName string
}
type Person struct {
    UserName string
    Emails []string
    Friends []*Friend
}

func main() {
    f1 := Friend{"minux.ma"}
    f2 := Friend{"tom"}
    t := template.New("fieldname example")
    t,_ = t.Parse(`Hello {{.UserName}}!
            {{range .Emails}}
             an email {{.}}
            {{end}}
            {{with .Friends}}
            {{range .}}
             myFriend name is {{.FName}}
            {{end}}
            {{end}}
        `)
    p := Person{UserName: "Jim",Emails: []string{"111313@qq.com","fefageaga@126.com"},Friends: []*Friend{&f1,&f2}}
    t.Execute(os.Stdout,p)
}
