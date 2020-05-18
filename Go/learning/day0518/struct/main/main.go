package main

import (
    "encoding/json"
    "fmt"
    "strconv"
)
type Monster struct {
    Name string `json:"name"`
    Age int `json:"age"`
    Skill string `json:"skill"`
    num int
}
func (m Monster) info(){
    fmt.Println(m)
}
func (m Monster) toString() (string) {
    jsonstr,_ :=json.Marshal(m)
    return string(jsonstr)
}
func (m Monster) run(speed int) {
    fmt.Println(m.Name, "run speed ",speed)
}
func (m *Monster) setName(name string) {
    m.Name= name
}
type Integer int
func (i Integer) toString() string {
    return strconv.Itoa(int(i))
}
func main() {
    type struct1 struct {
        name string
        age int
    }
    var s1 = struct1{
        name: "S1",
        age: 12,
    }
    fmt.Println(s1)
    var s2 = struct {
        name string
        age int
    }{
        name: "s2",
    }
    fmt.Println(s2)
    s1 = s2
    fmt.Println(s1)
    s3 := struct1{
        "s3",
         33,
    }
    s2 = s3
    fmt.Println(s2)
    var s4 *struct1 = new(struct1)
    fmt.Println(s4)
    fmt.Printf("s4 type: %T\n",s4)
    (*s4).age= 21
    (*s4).name = "s4"
    fmt.Println(*s4)
    fmt.Printf("*s4 type: %T\n",*s4)
    
    s5 := &struct1{
        name: "s5",
        age:  0,
    }
    s6 := s5
    fmt.Println("s5",s5)
    s6.name="s6"
    fmt.Println("s5",s5)
    fmt.Println("s6",s6)
    
    
    m1 := Monster{
        Name:  "M1",
        Age:   0,
        Skill: "kill",
    }
    jsonstr,err := json.Marshal(m1)
    if nil != err{
        fmt.Println("解析错误",err)
    }else{
        fmt.Printf("jsonStr = %q \n",jsonstr)
    }
    m2 := Monster{}
    err = json.Unmarshal(jsonstr,&m2)
    if nil != err {
        fmt.Println("解析错误",err)
    }else{
        fmt.Println(m2)
    }
   
    m4 := Monster{
        Name:  "m4",
        Age:   0,
        Skill: "cry",
    }
    m4.info()
    m2.info()
    m2.run(2)
    m2.setName("setName")
    m2.info()
    fmt.Println(m2.toString())
    num := m2.num
    fmt.Println(num)
    
    i := Integer(1)
    fmt.Println(i.toString())
    
}
