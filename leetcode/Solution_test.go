package leetcode

import (
    "fmt"
    "testing"
)
var list *ListNode = &ListNode{
    Val: 1,
    Next: &ListNode{
        Val:  2,
        Next: &ListNode{
            Val:  3,
            Next: &ListNode{
                Val:  4,
                Next: nil,
            },
        },
    },
}
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
func TestReverseList(t *testing.T) {
    //fmt.Println(ReverseList(list))
    
    fmt.Println(ReverseList2(list))
}
func TestIsPalindrome(t *testing.T){
    l1 := &ListNode{
        Val:  1,
        Next: &ListNode{
            Val:  2,
            Next: nil,
        },
    }
    fmt.Println(isPalindrome(l1))
    l1.Next = &ListNode{
        Val:  2,
        Next: &ListNode{
            Val: 1,
            Next: nil,
        },
    }
    fmt.Println(isPalindrome(l1))
}


func TestGetDecimalValue(t *testing.T) {
    l1 := &ListNode{
        Val:  1,
        Next: &ListNode{
            Val:  0,
            Next: &ListNode{
                Val:  1,
                Next: nil,
            },
        },
    }
    fmt.Println(getDecimalValue(l1))
}
func TestRemoveDuplicateNodes2(t *testing.T)   {
    l1 := &ListNode{
        Val:  1,
        Next: &ListNode{
            Val:  2,
            Next: &ListNode{
                Val:  2,
                Next: &ListNode{
                    Val:  1,
                    Next: nil,
                },
            },
        },
    }
    //head := removeDuplicateNodes2(l1)
    head := removeDuplicateNodes3(l1)
    fmt.Println(head)
}