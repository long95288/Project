package leetcode

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}


/**

 */
func hasCycle(head *ListNode) bool {
    cur := head
    m := make(map[*ListNode]interface{})
    for cur != nil{
        if _,ok := m[cur];ok {
            return true
        }
        m[cur] = nil
        cur = cur.Next
    }
    return false
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]interface{})
    if headA == nil || headB == nil {
        return nil
    }
    cur := headA
    for cur != nil {
        m[cur] = true
        cur = cur.Next
    }
    // 遍历B的时候,判断
    cur = headB
    for cur != nil {
        if _,ok := m[cur];ok{
            return cur
        }
        cur = cur.Next
    }
    return nil
}
// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
    length := 0
    cur := head
    if head == nil {
        return head
    }
    for cur != nil {
        length ++
        cur = cur.Next
    }
    cur = head
    if length % 2 == 0{
        // 偶数 = (length+2)/2
        length = (length + 2)/2
    }else {
        // 奇数 = (length+1)/2
        length = (length + 1)/2
    }
    for ;length>1;length-- {
        cur = cur.Next
    }
    return cur
}
/**
锯齿状数据,只能递减,每次减1。步数最少。
奇数位先减,操作的步数。
偶数位再减,操作的步数。
比较那个走得少。
 */
func MovesToMakeZigzag(nums []int) int {
    /**
    锯齿状数据,只能递减,每次减1。步数最少。
    奇数位先减,操作的步数。
    偶数位再减,操作的步数。
    比较那个走得少。
    */
    oddStep := 0
    eventStep := 0
    length := len(nums)
    // 奇位数先减 1,3,5,7
    for i:=0;i < length; i = i +2{
        tmp := nums[i]
        // 处在头部,只比较右边的数据
        if i == 0 {
            if nums[i + 1] > tmp {
                // A < B 高过，不用管
                continue
            }else {
                // A >= B,下拉
                oddStep = oddStep + (tmp - nums[i+1] + 1)
                continue
            }
        }
        // 处在尾部
        if i == length -1 {
            // B > C
            if nums[i-1] > tmp {
                continue
            }else {
                // B < C
                oddStep = oddStep + (tmp - nums[i-1] + 1)
                continue
            }
        }
        
        // 处在中间,比两边。低过最低的便可
        if i > 0 && i < length-1 {
            // 比较两边,
            if nums[i-1] > tmp && tmp < nums[i+1] {
                continue
            }
            if nums[i-1] > nums[i+1] {
                // A > C 低过右边便可
                if nums[i + 1] > tmp {
                    continue
                }else {
                    oddStep = oddStep + (tmp - nums[i+1] + 1)
                    continue
                }
            }else{
                // A < C 低过左边便可
                if nums[i - 1] > tmp {
                    continue
                }else {
                    oddStep = oddStep + (tmp - nums[i-1] + 1)
                    continue
                }
            }
        }
    }
    
    // 偶数位先减 2,4,6,8
    for i:=1;i < length;i= i + 2{
        tmp := nums[i]
        // 尾部
        if i == length -1 {
            // B > C
            if nums[i-1] > tmp {
                continue
            }else {
                // B < C
                eventStep = eventStep + (tmp - nums[i-1] + 1)
                continue
            }
        }
        // 中间
        if i > 0 && i < length-1 {
            // 比较两边,
            if nums[i-1] > tmp && tmp < nums[i+1] {
                continue
            }
            if nums[i-1] > nums[i+1] {
                // A > C 低过右边便可
                if nums[i + 1] > tmp {
                    continue
                }else {
                    eventStep = eventStep + (tmp - nums[i+1] + 1)
                    continue
                }
            }else{
                // A < C 低过左边便可
                if nums[i - 1] > tmp {
                    continue
                }else {
                    eventStep = eventStep + (tmp - nums[i-1] + 1)
                    continue
                }
            }
        }
    }
    
    if oddStep > eventStep {
        return eventStep
    }
    return oddStep
}
func ReverseList(head *ListNode) *ListNode {
   var tmp *ListNode = nil
   cur := head
   for cur != nil {
       tmpNode := &ListNode{
           Val: cur.Val,
       }
       tmpNode.Next = tmp
       tmp = tmpNode
       cur = cur.Next
   }
   return tmp
}
// 解法2
func ReverseList2(head *ListNode) *ListNode {
    var tmp *ListNode = nil
    cur := head
    for cur != nil {
        tmpNode := cur
        cur = cur.Next
        // 断开链表
        tmpNode.Next = nil
        // 重新连接
        tmpNode.Next = tmp
        tmp = tmpNode
    }
    return tmp
}
// 判断链表是否是回文数
func isPalindrome(head *ListNode) bool {
    // 复制链表
    cp := []int{}
    cur := head
    for cur != nil {
        cp = append(cp,cur.Val)
        cur = cur.Next
    }
    // 逆序比较
    cur = head
    index := len(cp) - 1
    index2 := 0
    for index != index2 && index2 < index {
        if cp[index] != cp[index2] {
            return false
        }
        index --
        index2 ++
    }
    //for cur != nil{
    //    if cur.Val != cp[index] {
    //        return false
    //    }
    //    index --
    //    cur = cur.Next
    //}
    return true
}
/**
二进制转10进制
 */
func getDecimalValue(head *ListNode) int {
    // 读出来
    //arr := []int{}
    cur := head
    length := 0
    for cur != nil {
        //arr = append(arr,cur.Val)
        cur = cur.Next
        length ++
    }
    // 按权重进行相加
    //length := len(arr)
    length --
    base := 1 << length
    result := 0
    for cur = head;cur!=nil;cur=cur.Next {
        result = result + cur.Val * base
        fmt.Println(result," ",base)
        base /= 2
    }
    return result
}
