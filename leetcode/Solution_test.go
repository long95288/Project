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
func TestPartition(t *testing.T)  {
    l1 := &ListNode{
        Val:  1,
        Next: &ListNode{
            Val:  4,
            Next: &ListNode{
                Val:  3,
                Next: &ListNode{
                    Val:  2,
                    Next: &ListNode{
                        Val:  5,
                        Next: &ListNode{
                            Val:  2,
                            Next: nil,
                        },
                    },
                },
            },
        },
    }
    partition(l1,3)
    printNode(l1)
}
func TestArrayPairSum(t *testing.T){
    arrayPairSum([]int{3,2,1,5,4})
}
func printNode(head *ListNode){
    for cur:=head;cur!=nil;cur=cur.Next{
        fmt.Println(cur.Val)
    }
}
func TestFindPairs(t *testing.T) {
    re := findPairs([]int{3, 1, 4, 1, 5},2)
    fmt.Println(re)
    re = findPairs([]int{1, 2, 3, 4, 5},1)
    fmt.Println(re)
    re = findPairs([]int{1, 3, 1, 5, 4},0)
    fmt.Println(re)
}
func TestFindUnsortedSubarray(t *testing.T){
    fmt.Printf("%d\n",findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
func TestCanPlaceFlowers(t *testing.T) {
    fmt.Printf("%t",canPlaceFlowers([]int{1,0,0,0,1,0,0},2))
}
func TestReorderList(t *testing.T)  {
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
                        Next: &ListNode{
                            Val:  6,
                            Next: nil,
                        },
                    },
                },
            },
        },
    }
    _ = l1
    l2 := &ListNode{
        Val:  1,
        Next: &ListNode{
            Val:  2,
            Next: &ListNode{
                Val:  3,
                Next: &ListNode{
                    Val: 4,
                    Next: nil,
                },
            },
        },
    }
    _ = l2
    reorderList(l1)
    printNode(l1)
}
func TestFindMaxAverage(t *testing.T)  {
    arr :=[]int{-6662,5432,-8558,-8935,8731,-3083,4115,9931,-4006,-3284,-3024,1714,-2825,-2374,-2750,-959,6516,9356,8040,-2169,-9490,-3068,6299,7823,-9767,5751,-7897,6680,-1293,-3486,-6785,6337,-9158,-4183,6240,-2846,-2588,-5458,-9576,-1501,-908,-5477,7596,-8863,-4088,7922,8231,-4928,7636,-3994,-243,-1327,8425,-3468,-4218,-364,4257,5690,1035,6217,8880,4127,-6299,-1831,2854,-4498,-6983,-677,2216,-1938,3348,4099,3591,9076,942,4571,-4200,7271,-6920,-1886,662,7844,3658,-6562,-2106,-296,-3280,8909,-8352,-9413,3513,1352,-8825}
    re := findMaxAverage(arr,90)
    fmt.Println("===================")
    arr2 := []int{9,7,3,5,6,2,0,8,1,9}
    re = findMaxAverage(arr2, 6)
    fmt.Println(re)
    
}
func TestString(t *testing.T)   {
    s := "Hello World"
    for i,c := range s{
        fmt.Println(i,c)
    }
}
func TestImageSmoother(t *testing.T)  {
    //[1,1,1],
    // [1,0,1],
    // [1,1,1]
    //向下舍入
    re := imageSmoother([][]int{{1,1,1},{1,0,1},{1,1,1}})
    fmt.Println(re)
    m := [][]int{
        {2, 3, 4},
        {5, 6, 7},
        {8, 9, 10},
        {11, 12, 13},
        {14,15,16},
    }
    re = imageSmoother(m)
    fmt.Println(re)
}
func TestCheckPossibility(t *testing.T)  {
   re := checkPossibility([]int{4,2,3})
   fmt.Println(re)
   re = checkPossibility([]int{4,2,1})
   fmt.Println(re)
   re = checkPossibility([]int{3,4,2,3})
   fmt.Println(re) // false
}
func TestSlice(t *testing.T) {
    p := []int{1,2,3,4,5,6}
    fmt.Println("size = ",len(p))
    p = p[:0]
    fmt.Println("size = ",len(p))
    s := Constructor1()
    s.Push(1)
    s.Pop()
    fmt.Println("size = ",s.Empty())
}
func TestNextGreaterElement(t *testing.T)  {
    nums11 := []int{4,1,2}
    nums21 := []int{1,3,4,2}
    ret := nextGreaterElement(nums11,nums21)
    fmt.Println(ret)
}
func TestCalPoints(t *testing.T)  {
    fmt.Println(calPoints([]string{"5","2","C","D","+"}))
    fmt.Println(calPoints([]string{"5","-2","4","C","D","9","+","+"}))
}
func TestBackspaceCompare(t *testing.T)  {
    fmt.Println(backspaceCompare("ab##","c#d#"))
}
func TestRemoveOuterParentheses(t *testing.T)  {
    fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
}
func TestAddBinary(t *testing.T)  {
    fmt.Println(addBinary("1","111"))
}
func TestIsPalindrome2(t *testing.T)  {
    fmt.Println(isPalindrome2("A man, a plan, a canal: Panama"))
    fmt.Println(isPalindrome2("race a car"))
    fmt.Println(isPalindrome2("0P"))
    fmt.Println(isPalindrome2("9,8"))
    fmt.Println(isPalindrome2("Zeus was deified, saw Suez."))
}
func TestLongestCommonPrefix(t *testing.T)  {
    fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
    // "dog","racecar","car"
    fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
}
func TestStrStr(t *testing.T)  {
    fmt.Println(strStr("hello","ll"))
    fmt.Println(strStr("aaaaaaa","bba"))
    fmt.Println(strStr("",""))
    fmt.Println(strStr("mississippi","issip"))
}
func TestMaxSlidingWindow(t *testing.T)  {
    fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7},3))
}
func TestIsHappy(t *testing.T)  {
    fmt.Println(isHappy(19))
}
func TestCountAndSay(t *testing.T)  {
    fmt.Println(countAndSay(11))
    fmt.Println("11131221133112132113212221")
}
func TestMySqrt(t *testing.T)  {
    fmt.Println(mySqrt(4))
    fmt.Println(mySqrt(8))
    fmt.Println(mySqrt(1085817232))
}
func TestCountPrimes(t *testing.T)  {
    fmt.Println(countPrimes(2))
    fmt.Println(countPrimes(10))
    fmt.Println(countPrimes2(2))
    fmt.Println(countPrimes2(10))
}
func TestSumOddLengthSubarrays(t *testing.T)  {
    fmt.Println(sumOddLengthSubarrays([]int{1,4,2,5,3}))
    fmt.Println(sumOddLengthSubarrays([]int{1,2}))
    fmt.Println(sumOddLengthSubarrays([]int{2}))
    fmt.Println(sumOddLengthSubarrays([]int{10,11,12}))
}
func TestMinSubarray(t *testing.T){
    // 1
    fmt.Println(minSubarray([]int{3,1,4,2},6))
    // 2
    fmt.Println(minSubarray([]int{6,3,5,2},9))
    // 0
    fmt.Println(minSubarray([]int{1,2,3},3))
    // -1
    fmt.Println(minSubarray([]int{1,2,3},7))
    
    // 1
    fmt.Println(minSubarray([]int{4,5,8,5,4},7))
    
    // -1
    fmt.Println(minSubarray([]int{4,4,2},7))
    
    // 0
    fmt.Println(minSubarray([]int{1},1))
    
}
func TestLongestUnivaluePath(t *testing.T){
    tree := TreeNode{
       Val:   5,
       Left:  &TreeNode{
           Val:   4,
           Left:  &TreeNode{
               Val:   1,
               Left:  nil,
               Right: nil,
           },
           Right: &TreeNode{
               Val:   1,
               Left:  nil,
           },
       },
       Right: &TreeNode{
           Val:   5,
           Left:  nil,
           Right: &TreeNode{
               Val:   5,
               Left:  nil,
           },
       },
    }
    fmt.Println(longestUnivaluePath(&tree))
    // [1,null,1,1,1,1,1,1]
    tree2 := &TreeNode{
       Val:   1,
       Left:  nil,
       Right: &TreeNode{
           Val:   1,
           Left:  &TreeNode{
               Val:   1,
               Left:  &TreeNode{
                   Val:   1,
                   Left:  nil,
                   Right: nil,
               },
               Right: &TreeNode{
                   Val:   1,
                   Left:  nil,
                   Right: nil,
               },
           },
           Right: &TreeNode{
               Val:   1,
               Left:  &TreeNode{
                   Val:   1,
                   Left:  nil,
                   Right: nil,
               },
               Right: nil,
           },
       },
    }
    fmt.Println(longestUnivaluePath(tree2))
    tree3 := &TreeNode{
        Val:   1,
        Left:  &TreeNode{
            Val:   2,
            Left:  &TreeNode{
                Val:   4,
                Left:  nil,
                Right: nil,
            },
            Right: &TreeNode{
                Val:   2,
                Left:  nil,
                Right: nil,
            },
        },
        Right: &TreeNode{
            Val:   3,
            Left:  nil,
            Right: nil,
        },
    }
    fmt.Println(longestUnivaluePath(tree3))
}

func TestFib(t *testing.T) {

}
func TestMissingNumber(t *testing.T){
    fmt.Println(missingNumber([]int{0,1,3}))
    
    fmt.Println(missingNumber([]int{0,1,2,3,4,5,6,7,9}))
    
    
    fmt.Println(missingNumber([]int{1}))
    
    fmt.Println(missingNumber([]int{0}))
}
func TestLevelOrderBottom(t *testing.T)  {
    tree := TreeNode{
        Val:   3,
        Left:  &TreeNode{
            Val:   9,
            Left:  nil,
            Right: nil,
        },
        Right: &TreeNode{
            Val:   20,
            Left:  &TreeNode{
                Val:   15,
                Left:  nil,
                Right: nil,
            },
            Right: &TreeNode{
                Val:   7,
                Left:  nil,
                Right: nil,
            },
        },
    }
    fmt.Println(levelOrderBottom(&tree))
}
func TestUncommonFromSentences(t *testing.T)  {
    fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
}

func TestIsIsomorphic(t *testing.T)  {
    fmt.Println(isIsomorphic("egg", "add"))
    fmt.Println(isIsomorphic("foo", "bar"))
    fmt.Println(isIsomorphic("paper", "title"))
    fmt.Println(isIsomorphic("", ""))
    fmt.Println(isIsomorphic("ab", "aa"))
}
func TestCanConstruct(t *testing.T)  {
    fmt.Println(canConstruct("a","b"))
    fmt.Println(canConstruct("aa", "ab"))
    fmt.Println(canConstruct("aa","aab"))
}
func TestFirstUniqChar(t *testing.T)  {
    fmt.Println(firstUniqChar("leetcode"))
    fmt.Println(firstUniqChar("loveleetcode"))
    fmt.Println(firstUniqChar("aadadaad"))
}
func TestReverseString(t *testing.T)  {
    a := []byte("Hello")
    reverseString(a)
    fmt.Println(string(a))
}
func TestReverseVowels(t *testing.T)  {
    fmt.Println(reverseVowels("leetcode"))
    fmt.Println(reverseVowels("hello"))
}
func TestAddStrings(t *testing.T)  {
    fmt.Println(addStrings("109","1"))
    fmt.Println(addStrings("999999999999999999999", "1"))
}