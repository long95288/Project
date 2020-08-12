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
    i = (-1) % 3
    fmt.Println(i)
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
func TestReversePrint(t *testing.T) {
    l1 := &ListNode{
        Val:  1,
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
    //fmt.Println(reversePrint(l1))
    fmt.Println(reversePrint2(l1))
}
func TestDeleteNode(t *testing.T)  {
    l1 := &ListNode{
        Val:  4,
        Next: &ListNode{
            Val:  5,
            Next: &ListNode{
                Val:  1,
                Next: &ListNode{
                    Val:  9,
                    Next: nil,
                },
            },
        },
    }
    deleteNode(l1)
    printNode(l1)
}
func TestDeleteNode2(t *testing.T) {
    l1 := &ListNode{
        Val:  1,
        Next: &ListNode{
            Val:  2,
            Next: &ListNode{
                Val:  2,
                Next: &ListNode{
                    Val:  5,
                    Next: nil,
                },
            },
        },
    }
    re := deleteDuplicates2(l1)
    printNode(re)
}
func TestReverseBetween(t *testing.T)  {
    l1 := &ListNode{
       Val:  1,
       Next: &ListNode{
           Val:  2,
           Next: &ListNode{
               Val:  3,
               Next: &ListNode{
                   Val:  4,
                   Next: &ListNode{
                       Val:  5,
                       Next: nil,
                   },
               },
           },
       },
    }
    printNode(reverseBetween(l1,1,5))
    l2 := &ListNode{
       Val:  3,
       Next: &ListNode{
           Val:  5,
           Next: nil,
       },
    }
    fmt.Println(".....")
    printNode(reverseBetween(l2,1,2))
}
func TestReverseBetween2(t *testing.T)  {
    l2 := &ListNode{
        Val:  3,
        Next: &ListNode{
            Val:  4,
            Next: &ListNode{
                Val:  5,
                Next: &ListNode{
                    Val:  6,
                    Next: nil,
                },
            },
        },
    }
    fmt.Println(".....")
    printNode(reverseBetween2(l2,2,3))
}
func TestAddTwoNumbers(t *testing.T)  {
    l1 := &ListNode{
        Val:  7,
        Next: &ListNode{
            Val:  2,
            Next: &ListNode{
                Val:  4,
                Next: &ListNode{
                    Val:  3,
                    Next: nil,
                },
            },
        },
    }
    l2 := &ListNode{
        Val:  5,
        Next: &ListNode{
            Val:  6,
            Next: &ListNode{
                Val:  4,
                Next: nil,
            },
        },
    }
    printNode(addTwoNumbers(l1,l2))
}

func TestGenerator(t *testing.T) {
    fmt.Println(generate(5))
}
func TestRotate(t *testing.T) {
    //rotate([]int{17},3)
    getRow(3)
}
func printNode(head *ListNode){
    for cur:=head;cur!=nil;cur=cur.Next{
        fmt.Println(cur.Val)
    }
}
