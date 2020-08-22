package test

import (
    "fmt"
    "strings"
    "testing"
)

func TestTrim(t *testing.T)  {
    fmt.Printf("%q",strings.TrimSpace(" Hello world    "))
    
}
