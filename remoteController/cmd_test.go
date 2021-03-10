package main

import (
    "fmt"
    "testing"
)

func TestGetScreenCapture(t *testing.T)  {
    out, err := GetScreenCapture()
    fmt.Printf("out size %d err : %v \n",len(out), err)
    out, err = resizeImage(out)
    fmt.Printf(" resize %d err : %v \n",len(out), err)
}