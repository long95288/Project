package _struct

import "encoding/json"

type Student struct {
    age int `json:"age"`
    name string `json:"name"`
}

func (s Student) toString()string {
    str,_ := json.Marshal(s)
    return string(str)
}


