package leetcode

import (
    "fmt"
    "testing"
)

func TestDivide(t *testing.T)  {
    fmt.Println(5/2)
    fmt.Println(6/2)
}
func TestAbs(t *testing.T) {
    i := -1
    fmt.Println(-i)
}
func TestMovesToMakeZigzag(t *testing.T)  {
    nums := []int{1,2,3}
    re :=  MovesToMakeZigzag(nums)
    fmt.Println(re)
    
    nums = []int{9,6,1,6,2}
    re = MovesToMakeZigzag(nums)
    fmt.Println(re)
}