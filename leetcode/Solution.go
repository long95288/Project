package leetcode

import (
    "fmt"
    "math"
    "sort"
    "strconv"
)

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
func getKthFromEnd(head *ListNode, k int) *ListNode {
    // 第一遍算长度
    length := 0
    cur := head
    for ;cur != nil;cur= cur.Next {
        length ++
    }
    // 算前进步数
    length = length - k
    cur = head
    for i:=1; i <=length; i++{
        cur = cur.Next
    }
    return cur
}
func getKthFromEnd2(head *ListNode, k int) *ListNode {
    fast,slow := head,head
    for i:=k;i>0;i--{
        fast = fast.Next
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
func removeDuplicateNodes(head *ListNode) *ListNode {
    // 向后读取,如果后面的数据和现在的一样,删除节点。不一样,指针前移
    if head == nil {
        return head
    }
    pre,cur := head,head.Next
    // 临时缓冲区
    buf := make(map[int]struct{})
    buf[pre.Val] = struct{}{}
    for cur != nil {
        if _,ok := buf[cur.Val];ok {
            // 相同,删除节点
            cur = cur.Next
            pre.Next = cur
        }else{
            // 不相同,继续
            buf[cur.Val] = struct{}{}
            pre = pre.Next
            cur = cur.Next
        }
    }
    return head
}
func removeDuplicateNodes2(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    pre,cur := head,head.Next
    for pre != nil {
        preTmp := pre
        for cur != nil {
            if cur.Val == pre.Val {
                cur = cur.Next
                preTmp.Next = cur
            }else{
                preTmp = preTmp.Next
                cur = cur.Next
            }
        }
        pre = pre.Next
        if pre == nil {
            break
        }
        cur = pre.Next
    }
    return head
}
func removeDuplicateNodes3(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    index := head
    for index != nil {
        cur := index
        for cur.Next != nil {
            if cur.Next.Val == index.Val {
                // 删除节点
                cur.Next = cur.Next.Next
            }else{
                cur = cur.Next
            }
        }
        index = index.Next
    }
    return head
}
func reversePrint(head *ListNode) []int {
    arr := []int{}
    for cur := head;cur != nil;cur = cur.Next{
        arr = append([]int{cur.Val},arr...)
    }
    return arr
}
func reversePrint2(head *ListNode) []int {
    length := 0
    for cur:=head;cur != nil;cur=cur.Next{
        length ++
    }
    arr := make([]int,length)
    for cur:=head;cur!=nil;cur=cur.Next{
        arr[length-1] = cur.Val
        length --
    }
    return arr
}
func deleteNode(node *ListNode) {
    // 快慢指针
    // 快指针比慢指针快1步，慢指针需要后退一步,再继续
    step := 2
    fast,slow := node,node
    for fast!=nil && fast.Next != nil {
        if step > 0 {
            // 停留一步
            step --
        }else{
            slow = slow.Next
        }
        
        fast = fast.Next.Next
    }
    // 删除中间节点
    slow.Next = slow.Next.Next
}
func deleteDuplicates(head *ListNode) *ListNode {
    // map
    mp := make(map[int]int)
    for cur:=head;cur!=nil;cur=cur.Next{
        mp[cur.Val] += 1
    }
    dump := &ListNode{}
    dump.Next = head
    dup := false
    for cur := dump;cur!= nil&&cur.Next != nil;{
        // 判断重复 值大于2
        if v := mp[cur.Next.Val]; !dup && v > 1 {
            dup = true
        }
        // 删除重复的节点
        if dup {
            if v := mp[cur.Next.Val];v > 0{
                mp[cur.Next.Val] -= 1
                if v := mp[cur.Next.Val];v == 0 {
                    dup = false
                }
                cur.Next = cur.Next.Next
            }
        }else{
            cur = cur.Next
        }
    }
    return dump.Next
}
func deleteDuplicates2(head *ListNode) *ListNode {
    //if head == nil || head.Next == nil {
    //    return head
    //}
    //// 小于三个个元素的
    //if head.Next.Next == nil {
    //    if head.Val == head.Next.Val{
    //        return nil
    //    }else{
    //        return head
    //    }
    //}
    //// 三个元素以上的
    //// 新链表
    //newLink := &ListNode{}
    //// 新链表的尾巴指针
    //newLinkTail := newLink
    //// 前一个值,后一个值
    //beforeVal := head.Val
    //for cur := head.Next;cur != nil;{
    //    // 中间节点
    //    if cur.Next == nil {
    //        // 最后一个了,只要比较前面就行
    //        if beforeVal != cur.Val {
    //            // 不同,加入新链表
    //            newLinkTail.Next = cur
    //            break
    //        }else {
    //            //相同 掠过
    //            cur = cur.Next
    //        }
    //    }else{
    //        // 不是最后一个
    //        if beforeVal != cur.Val && cur.Val != cur.Next.Val {
    //            // 都不相同,加入节点
    //            newLinkTail.Next = cur
    //            // 更新值
    //            beforeVal = cur.Val
    //            cur = cur.Next
    //
    //            newLinkTail= newLinkTail.Next
    //            newLinkTail.Next = nil
    //        }else if beforeVal == cur.Val{
    //            cur = cur.Next
    //        }
    //    }
    //}
    //return newLink.Next
    return nil
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    
    // 先将待翻转的数据存下来,存完之后就翻转
    // 使用array存储
    // 哑节点
    if head.Next == nil {
        return head
    }
    dummy := &ListNode{}
    dummy.Next = head
    pre,cur := dummy,head
    arr := []*ListNode{}
    sub := n - m
    step := 2
    i := 0
    for {
        if step == 2 {
            // 分析阶段
            if i < m -1 {
                // 定位
                pre = cur
                cur = cur.Next
                i ++
            }else if i >= m -1 && i <= n -1{
                // 赋值
                arr = append(arr,cur)
                cur = cur.Next
                i ++
            }else{
                step --
                continue
            }
        }else if step == 1{
            // 反转阶段
            if sub >= 0 {
                tmp := arr[sub]
                tmp.Next = nil
                pre.Next = tmp
                pre = tmp
                sub --
            }else{
                // 翻转完成,链表连接
                pre.Next = cur
                step --
            }
        }else{
            break
        }
    }
    return dummy.Next
}
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
    // 使用多指针，边走边翻转
    dummy := &ListNode{}
    dummy.Next = head
    // 新生成链表的尾巴
    dTail := dummy
    // 旧的数据
    var sHead *ListNode = nil
    var sTail *ListNode = nil
    i := 0
    cur := head
    for {
        if i < m -1 {
            dTail = cur
            cur = cur.Next
            i ++
        }else if i == m-1{
            // 切割点
            sTail = cur
            sHead = cur
            cur = cur.Next
            i ++
        }else if i > m -1 && i <= n -1 {
            // 翻转点
            tmp := cur
            cur = cur.Next
            tmp.Next = sHead
            sHead = tmp
            i ++
        }else if i == n {
            // 重新连接点
            dTail.Next = sHead
            sTail.Next = cur
            break
        }else{
            break
        }
    }
    
    return dummy.Next
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stack1 := []int{}
    stack2 := []int{}
    // 数据入栈
    for cur:=l1;cur!= nil;cur=cur.Next{
        stack1 = append(stack1,cur.Val)
    }
    for cur:=l2;cur!=nil;cur=cur.Next{
        stack2 = append(stack2,cur.Val)
    }
    // 出栈相加
    dummy := &ListNode{}
    top1,top2 := len(stack1),len(stack2)
    carray := 0
    for top1 > 0 || top2 > 0{
        a,b := 0,0
        if top1 == 0 {
            a = 0
        }else{
            a = stack1[top1 - 1]
            top1 --
        }
        if top2 == 0 {
            b = 0
        }else{
            b = stack2[top2 -1]
            top2 --
        }
        sum := a + b + carray
        carray = sum / 10
        val := sum % 10
        node := &ListNode{
            Val:val,
            Next:nil,
        }
        node.Next = dummy.Next
        dummy.Next = node
    }
    if carray > 0 {
        cNode := &ListNode{
            Val:1,
            Next:nil,
        }
        cNode.Next = dummy.Next
        dummy.Next = cNode
    }
    return dummy.Next
}
func searchInsert(nums []int, target int) int {
    i := 0
    for ;i < len(nums);i++{
        if nums[i] == target {
            return i
        }
        if i + 1 < len(nums)  {
            if target > nums[i] && target < nums[i + 1]{
                return i + 1
            }
        }
    }
    return i
}
func merge(A []int, m int, B []int, n int)  {
    // 从后向前移动,每次寻找最大值填入A
    m,n = m-1,n-1
    index := m + n -1
    for m >=0 || n >= 0 {
        if m < 0 {
            // A已经选取结束了,把B放入便可
            A[index] = B[n]
            index --
            n --
            continue
        }
        if n < 0 {
            A[index] = A[m]
            index --
            m --
            continue
        }
        if A[m] >= B[n] {
            A[index] = A[m]
            m --
            index --
        }else {
            A[index] = B[n]
            n--
            index --
        }
    }
}
func generate(numRows int) [][]int {
   reNums := make([][]int,numRows)
   for i:=0;i<numRows;i++{
       if i == 0 {
           reNums[i] = []int{1}
       }else{
           subNum := make([]int,i+1)
           for j:=0;j < i + 1;j++ {
                pre1,pre2 := 0,0
                if j-1 <0 {
                    pre1 = 0
                }else{
                    pre1 = reNums[i-1][j-1]
                }
                
                if j >= i {
                    pre2 = 0
                }else{
                    pre2 = reNums[i-1][j]
                }
                subNum[j] = pre1 + pre2
           }
           reNums[i] = subNum
       }
   }
   return reNums
}
func reverse(nums []int, start int, end int) {
    for start < end{
        tmp := nums[start]
        nums[start] = nums[end]
        nums[end] = tmp
        start ++
        end --
    }
}
func rotate(nums []int, k int)  {
    // 先整体翻转,然后前k个数据翻转，然后再length - K个数据翻转
    length := len(nums)
    k = k % length
    // 1.全部翻转
    reverse(nums,0,length-1)
    // 2.翻转前k个
    reverse(nums,0,k-1)
    // 3.翻转后n-k个
    reverse(nums,k,length-1)
}

func factorial(n int64) int64{
    if n == 0 {
        return 1
    }
    return factorial(n-1)*n
}
func combination(n1, m1 int) int64 {
    n := int64(n1)
    m := int64(m1)
    if m == 0 {
        return 1
    }
    fn := factorial(n)
    fm := factorial(m)
    fnm := factorial(n-m)
    return (fn/fm)*(fn/fnm)
}
func getRow(rowIndex int) []int {
    // 只要计算一半便可
    half := rowIndex/2
    nums := make([]int,rowIndex + 1)
    for i:=0;i<=half;i++{
        nums[i] = int(combination(rowIndex,i))
    }
    for i:= half+1;i < rowIndex + 1;i++{
        nums[i] = nums[rowIndex - i]
    }
    return nums
}
// node插入到head的对应的位置
func insertNode(head *ListNode,node *ListNode){
    for cur := head;cur != nil;cur = cur.Next {
        if cur.Next != nil {
            if node.Val >= cur.Val && node.Val < cur.Next.Val {
                node.Next = cur.Next
                cur.Next = node
                break
            }
        }else{
            node.Next = nil
            cur.Next = node
            break
        }
    }
}
func partitionOld(head *ListNode, x int) *ListNode {
    // 构建两条链表,分别存放符合要求的。最后连接起来
    first,second := &ListNode{},&ListNode{}
    firstTail := first
    for cur:=head;cur!=nil;{
        if cur.Val <= x {
            // 搜索插入first表
            tmp := cur
            cur = cur.Next
            insertNode(first,tmp)
        }else{
            // 搜索插入second表
            tmp := cur
            cur = cur.Next
            insertNode(second,tmp)
        }
    }
    for cur:=first;cur!=nil;cur=cur.Next{
        firstTail = cur
    }
    // 重新连结
    firstTail.Next= second.Next
    return first.Next
}
func partition(head *ListNode, x int) *ListNode {
    first,second := &ListNode{},&ListNode{}
    firstTail,secondTail := first,second
    // 进行分割
    for cur := head;cur!= nil;{
        if cur.Val < x {
            // 添加到first链表
            tmp := cur
            cur = cur.Next
            tmp.Next = nil
            firstTail.Next = tmp
            firstTail = tmp
        } else {
            // 添加到second链表
            tmp := cur
            cur = cur.Next
            tmp.Next = nil
            secondTail.Next = tmp
            secondTail = tmp
        }
    }
    
    // 两个链表拼接
    firstTail.Next = second.Next
    return first.Next
}
func arrayPairSum(nums []int) int {
    // 排序
    sort.Ints(nums)
    // 奇数位相加
    sum := 0
    for i:=0;i<len(nums);i++{
        if i % 2 == 0 {
            sum += nums[i]
        }
    }
    return sum
}
func findPairs(nums []int, k int) int {
    if k < 0 {
        return 0
    }
    numsHas := make(map[int]bool)
    diffHas := make(map[int]bool)
    
    for _, num := range nums {
        if numsHas[num - k] {
            diffHas[num - k] = true
        }
        if numsHas[num + k] {
            diffHas[num] = true
        }
        numsHas[num] = true
    }
    return len(diffHas)
}
func matrixReshape(nums [][]int, r int, c int) [][]int {
    row := len(nums)
    column := len(nums[0])
    if r * c != row * column {
        return nums
    }
    list := []int{}
    for _,rowItem := range nums {
        for _,columnItem := range rowItem {
            list = append(list,columnItem)
        }
    }
    count := 0
    reNums := make([][]int,r)
    for i:=0;i<r;i++ {
        colNum := make([]int,c)
        for j:=0;j<c;j++{
            colNum[j] = list[count]
            count ++
        }
        reNums[i] = colNum
    }
    return reNums
}

func findUnsortedSubarray(nums []int) int {
    newNum := make([]int,len(nums))
    for i,num := range nums {
        newNum[i] = num
    }
    sort.Ints(newNum)
    fmt.Println(nums)
    fmt.Println(newNum)
    head := 0
    for i,num := range nums {
        if num != newNum[i] {
            head = i
            break
        }
    }
    tail := len(nums) -1
    for i:=tail;i>=0;i--{
        if nums[i] != newNum[i] {
            tail = i
            break
        }
    }
    return tail - head + 1
}
func canPlaceFlowers(flowerbed []int, n int) bool {
    f := append([]int{0},append(flowerbed,0)...)
    count := 0
    for i := 1;i <len(f)-1;i++ {
        if count == n {
            return true
        }
        if f[i-1] == 0 && f[i] == 0 && f[i+1] == 0 {
            count ++
            f[i] = 1
        }
    }
    return count == n
}
/**
给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func reorderList(head *ListNode)  {
    // 分割 倒序 重连
    // 寻找中间节点
    slow,fast := head,head
    for ;fast != nil && fast.Next != nil;{
        fast = fast.Next.Next
        slow =slow.Next
    }
    // 从中间分割
    cur := slow.Next
    slow.Next = nil
    // 后面的链表倒置
    var newListHead *ListNode = nil
    for ; cur != nil;{
        tmp := cur
        cur = cur.Next
        tmp.Next = nil
        tmp.Next = newListHead
        newListHead = tmp
    }
    // 连接
    cur = head
    for cur2 := newListHead;cur2 != nil;{
        temp := cur2
        cur2 = cur2.Next
        temp.Next = cur.Next
        cur.Next = temp
        cur = cur.Next.Next
    }
}
func length(head *ListNode) int  {
    length :=0
    for cur:=head;cur!=nil;cur = cur.Next{
        length++
    }
    return length
}

func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    max := math.MinInt64
    for i:=0;i<k;i++{
        sum += nums[i]
    }
    if sum > max{
        max = sum
    }
    for i := k;i < len(nums);i++{
        sum += nums[i]
        sum = sum - nums[i-k]
        if sum > max {
            max = sum
        }
    }
    return float64(max) / float64(k)
}
func isValid(s string) bool {
    stack := make([]int32,len(s))
    top := -1
    for _,c := range s {
       if top == -1 {
           top ++
           stack[top] = c
       }else{
            // '(' ')' '{' '}' '[' ']'
           if (stack[top] == '(' && c == ')') || (stack[top] == '{' && c == '}') || (stack[top] == '[' && c == ']') {
               top --
           }else{
               top ++
               stack[top] = c
           }
       }
    }
    return top < 0
}
type MinStack struct {
    Val []int
    Min []int
    TopVal int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    stack := MinStack{
        Val : []int{},
        Min : []int{},
        TopVal : 0,
    }
    stack.Val = append(stack.Val,-1)
    stack.Min = append(stack.Min,math.MinInt32)
    return stack
}


func (this *MinStack) Push(x int)  {
    if this.TopVal == 0{
        this.Min = append(this.Min,x)
    }else if x < this.Min[this.TopVal] {
        this.Min = append(this.Min,x)
    }else{
        this.Min = append(this.Min,math.MinInt32)
        this.Min[this.TopVal + 1] = this.Min[this.TopVal]
    }
    this.TopVal += 1
    this.Val = append(this.Val,x)
}

func (this *MinStack) Pop()  {
    this.TopVal -= 1
    this.Val = this.Val[:this.TopVal + 1]
    this.Min = this.Min[:this.TopVal + 1]
}


func (this *MinStack) Top() int {
    return this.Val[this.TopVal]
}


func (this *MinStack) GetMin() int {
    return this.Min[this.TopVal]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func getNum(c int32) int {
    switch c {
    case 'I':
        return 1
    case 'V':
        return 5
    case 'X':
        return 10
    case 'L':
        return 50
    case 'C':
        return 100
    case 'D':
        return 500
    case 'M':
        return 1000
    default:
        return 0
    }
}
func romanToInt(s string) int {
    sum := 0
    pre := 0
    for i,v := range s{
        if i == 0{
            pre = getNum(v)
        }else{
            num := getNum(v)
            if pre < num {
                sum -= pre
            }else{
                sum += pre
            }
            pre = num
        }
    }
    sum += pre
    return sum
}
/**
图片平滑器
 */
func imageSmoother(M [][]int) [][]int {
    row := len(M)
    column := len(M[0])
    re := make([][]int,row)
    for i:=0;i<row;i++{
        re[i] = make([]int,column)
        for j := 0; j < column; j++ {
            count := 0
            sum := 0
            // 获得周围8格的数据
            for i1:=-1; i1 < 2; i1 ++{
                for j1:=-1;j1 < 2;j1++{
                    if i1 + i >= 0 && i1 + i < row  && j1 + j >= 0 && j1+ j < column{
                        sum += M[i1+i][j1+j]
                        count ++
                    }
                }
            }
            // 除以平均数
             re[i][j] = int(math.Floor(float64(sum) / float64(count)))
        }
    }
    return re
}
func checkPossibility(nums []int) bool {
    if len(nums) <=1 {
        return true
    }
    count := 0
    for i:=1;i<len(nums) && count < 2;i ++{
        if nums[i-1] <= nums[i]{
            continue
        }
        count ++
        if i-2 >=0 &&nums[i-2] > nums[i] {
            nums[i] = nums[i-1]
        }else {
            nums[i -1] = nums[i]
        }
    }
    return count <= 1
}

type MyStack struct {
    top int
    p1 []int
}


/** Initialize your data structure here. */
func Constructor1() MyStack {
    return MyStack{p1: []int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.p1 = append(this.p1,x)
    this.top = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    top := this.top
    this.p1 = this.p1[:len(this.p1)-1]
    if len(this.p1) <= 0{
        this.top = -1
    }else {
        this.top = this.p1[len(this.p1)-1]
    }
    return top
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.p1) <= 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    
    // 先计算num2每个元素更大的
    indexMap := make(map[int]int)
    for i:=0;i<len(nums2);i++{
        search := false
        for j:=i+1;j<len(nums2);j++{
            if nums2[j] > nums2[i] {
                indexMap[nums2[i]] = nums2[j]
                search = true
                break
            }
        }
        if !search{
            indexMap[nums2[i]] = -1
        }
    }
    ret := []int{}
    for _,v := range nums1{
       mv,_ := indexMap[v]
       ret = append(ret,mv)
    }
    return ret
}
/**
数字入栈
+ 两个数出栈相加并入栈
D 出栈*2入栈
C 最上面的出栈
 */
func calPoints(ops []string) int {
    stack := make([]int,len(ops))
    re := 0
    top := -1
    for _,v := range ops{
        //fmt.Println(v)
        switch v {
        case "C":
            re -= stack[top]
            top --
            
        case "D":
            tmp := 2 * stack[top]
            top ++
            stack[top] = tmp
            re += tmp
        case "+":
            tmp1 := stack[top]
            tmp2 := 0
            if top > 0 {
                top --
                tmp2 = stack[top]
                top ++
            }
            top ++
            stack[top] = tmp1 + tmp2
            re = re + tmp1 + tmp2
        default:
            d,_ := strconv.Atoi(v)
            top ++
            stack[top] = d
            re += d
        }
    }
    return re
}
func backspaceCompare(S string, T string) bool {
    stack1 := make([]rune,len(S))
    stack2 := make([]rune,len(T))
    top1,top2 := -1,-1
    for _,v := range S {
        if '#' == v{
            if top1 >= 0{
                top1 --
            }
        }else{
            top1 ++
            stack1[top1] = v
        }
    }
    for _,v := range T {
        if '#' == v {
            if top2 >= 0 {
                top2 --
            }
        }else{
            top2 ++
            stack2[top2] = v
        }
    }
    if top1 == top2 {
        for ;top1 >= 0;top1--{
            if stack1[top1] != stack2[top1] {
                return false
            }
        }
        return true
    }
    return false
}
func removeOuterParentheses(S string) string {
    stack := make([]rune,len(S))
    s := make([]rune,len(S))
    index := -1
    top := -1
    //re := ""
    for _,v := range S{
        if v == '('{
            if  top >= 0 {
                // 不是最外层括号,需要打印出
                //re += "("
                index ++
                s[index] = '('
                //是最外层括号,只要进栈
            }
            top ++
            stack[top] = v
        }else {
            if top > 0 {
                // 打印非最外层
                //re += ")"
                index ++
                s[index] = ')'
            }
            top --
        }
    }
    return string(s[:index+1])
}
func removeDuplicates(S string) string {
    // 入栈,相同出栈
    stack := make([]rune,len(S))
    top := -1
    for _,v := range S{
        if top < 0{
            top++
            stack[top] = v
        }else{
            if stack[top] == v {
                top --
            }else{
                top++
                stack[top] = v
            }
        }
    }
    return string(stack[:top+1])
}
func addBinary(a string, b string) string {
    aBit := make([]int,len(a))
    bBit := make([]int,len(b))
    stack := make([]int,len(a)+len(b))
    top := -1
    for i, v := range a {
       d,_ := strconv.Atoi(string(v))
       aBit[i] = d
    }
    for i, v := range b {
        d,_ := strconv.Atoi(string(v))
        bBit[i] = d
    }
    i,j := len(a)-1,len(b)-1
    
    tmpa := aBit[0]
    tmpb := bBit[0]
    carry := 0
    for i >= 0 || j >=0 {
        if i < 0 {
            tmpa = 0
        }else{
            tmpa = aBit[i]
        }
        if j < 0 {
            tmpb = 0
        }else {
            tmpb = bBit[j]
        }
        sum := tmpa + tmpb + carry
        carry = sum / 2
        top ++
        stack[top] = sum % 2
        i --
        j --
    }
    if carry > 0 {
        top ++
        stack[top] = carry
    }
    ret := make([]rune,top+1)
    stack_size := top
    for top >= 0 {
        ret[stack_size - top] = rune(48 + stack[top])
        top --
    }
    return string(ret[:stack_size + 1])
}

func isPalindrome2(s string) bool {
    if "" == s {
        return true
    }
    stack := make([]rune,len(s))
    top := -1
    
    tmp_s := make([]rune,len(s))
    index_tmp := -1
    for _,v := range s {
        if (v >= '0' && v <= '9') || (v >= 'a' && v <='z') || (v >='A' && v <= 'Z') {
            top ++
            stack[top] = v
            
            index_tmp ++
            tmp_s[index_tmp] = v
        }
    }
    
    for i:=0;i<= top;i++{
        v := tmp_s[i]
        tmp := stack[top - i]
        if (v >= 'a' && v <='z') || (v >='A' && v <= 'Z') {
            sub := math.Abs(float64(tmp - v))
            if sub == 0 || sub == 32{
                continue
            }else {
                return false
            }
        }else{
            if v == tmp {
                continue
            }
            return false
        }
    }
    return true
}

func longestCommonPrefix(strs []string) string {
    prefix := ""
    if len(strs) == 1 {
        return strs[0]
    }
    if len(strs) == 2 {
        first := strs[0]
        second := strs[1][:]
        for i,v := range first {
            if i < len(second){
                if uint8(v) == second[i] {
                    prefix += string(v)
                }else{
                    break
                }
            }else{
                break
            }
        }
        return prefix
    }
    prefix = longestCommonPrefix(strs[:2])
    return longestCommonPrefix(append([]string{prefix},strs[2:]...))
}

func strStr(haystack string, needle string) int {
    if "" == needle {
        return 0
    }
    // 思路,for循环,比较needle。完全匹配才可以
    needle_ := needle[:]
    haystack_ := haystack[:]
    for i:=0;i<len(haystack);i++{
        v := haystack_[i]
        if v == needle_[0]{
            // 开始进行匹配
            j := i
            index_needle := 0
            for index_needle < len(needle) && j < len(haystack) {
                if index_needle == len(needle) -1 && needle_[index_needle] == haystack_[j] {
                    return i
                }
                if haystack_[j] != needle_[index_needle]  {
                    break
                }else{
                    j ++
                    index_needle ++
                }
            }
        }
    }
    return -1
}
func Max(nums []int) int {
    max := math.MinInt32
    for _,v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) <= 0 {
        return []int{0}
    }
    if len(nums) <= k {
        return []int{Max(nums)}
    }
    slideWindows := make([]int,k)
    baseMax := math.MinInt32
    ret := make([]int, len(nums) - k + 1)
    retIndex := 0
    for i:=0;i<k && i <len(nums);i++{
        if nums[i] > baseMax {
            baseMax = nums[i]
        }
        slideWindows[i] = nums[i]
    }
    ret[retIndex] = baseMax
    for i:=k;i<len(nums);i++{
        slideWindows = append(slideWindows[1:],nums[i])
        baseMax = Max(slideWindows)
        retIndex ++
        ret[retIndex] = baseMax
    }
    return ret
}
func isHappy(n int) bool {
    strn := strconv.FormatInt(int64(n),10)
    result := 0
    for _,v := range strn{
        tmp,_ := strconv.Atoi(string(v))
        result += tmp * tmp
    }
    if result == 1 {
        return true
    }else if result == 4{
        return false
    }
    return isHappy(result)
}
func countAndSay(n int) string {
    if n <= 1 {
        return "1"
    }
    result := ""
    tmpStr := countAndSay(n - 1)
    tmp := tmpStr[:]
    count := 0
    tmpChar := tmp[0]
    // 数数
    for i := 0; i < len(tmp); i++{
        if tmpChar == tmp[i] {
            count ++
        }else{
            // 出现不同
            result += strconv.Itoa(count) + string(tmpChar)
            tmpChar = tmp[i]
            count = 1
        }
    }
    if count > 0 {
        result += strconv.Itoa(count) + string(tmpChar)
    }
    return result
}
func mySqrt(x int) int {
    if x == 1 {
        return 1
    }
    half := x/2
    start := 1
    end := half
    mid :=  (start + end) / 2
    for i:=1;i <= 1000;i++{
        if start * start == x {
            return start
        }
        if end * end == x {
            return end
        }
        if start*start < x && end * end > x && end - start == 1 {
            return start
        }
        mid = (start + end) / 2
        if mid * mid < x {
            start = mid
        }else {
            end = mid
        }
    }
    return half
}
func isPrimes(n int) bool{
    for i:=2; i<n; i++{
        if (float64(n)/float64(i) - float64(n/i)) == 0 {
            return false
        }
    }
    return true
}
func countPrimes(n int) int {
    count := 0
    if n < 2 {
        return 0
    }
    if n == 2 {
        return 1
    }
    count ++
    for i := 3;i<n;i++{
        if isPrimes(i) {
            count++
        }
    }
    return count
}
func countPrimes2(n int) int {
    count := 0
    if n <= 2 {
        return 0
    }
    count ++
    arr := make([]bool,n)
    for i := 3;i < n;i += 2{
        if !arr[i] {
            for j:=3;i*j <n;j+= 2{
                arr[i*j] = true
            }
            count ++
        }
    }
    return count
}
func sumArray(arr []int) int {
    result := 0
    for _,v := range arr{
        result += v
    }
    return result
}
func sumOddLengthSubarrays(arr []int) int {
    result := 0
    if len(arr) <= 0 {
        return 0
    }
    // 设置滑动窗口
    slide_window := make([]int,1)
    slide_index := 0
    slide_size := 1
    for ;slide_size <= len(arr);slide_size += 2{
        slide_index = 0
        slide_window = make([]int,slide_size)
        for _,v := range arr {
            // 先填充滑动窗口
            if slide_index <= slide_size -1 {
                slide_window[slide_index] = v
                slide_index ++
            }
            if slide_index == slide_size {
                // 填满之后计算结果
                result += sumArray(slide_window)
                slide_window = append(slide_window[1:],0)
                slide_index -= 1
            }
        }
    }
    return result
}

func minSubarray(nums []int, p int) int {
    // 先进行简化数组
    for i:=0;i < len(nums);i++{
        nums[i] = nums[i] % p
    }
    sum := 0
    setSum := false
    for _,v := range nums{
        sum += v
        if sum >= p {
            setSum = true
        }
        if sum % p == 0 {
            sum = 0
            setSum = true
        }
        
    }
    if !setSum {
        return -1
    }
    
    if sum == 0 {
        return 0
    }
    sum = sum % p
    // 最后剩余的值看能不能通过寻找子数组找出来
    removeSize := -1

    subSum := 0
    i := 0
    for j:=i;j<len(nums);j++{
        subSum += nums[j]
        subSum = subSum % p
        if subSum == 0 {
            i = j
        }
        if sum == nums[j] {
            return 1
        }else if sum == subSum {
             tmp := j - i
             if removeSize == -1 {
                 removeSize = tmp
             }else if tmp < removeSize {
                 removeSize = tmp
             }
        }
    }
    if removeSize >= len(nums) -1 {
        return -1
    }
    return removeSize
}
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Root2Leaves(root *TreeNode) int {
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    }
    left := Root2Leaves(root.Left)
    leftEqual := false
    right := Root2Leaves(root.Right)
    rightEqual := false
    
    if root.Left != nil && root.Left.Val == root.Val {
        left += 1
        leftEqual = true
    }
    if root.Right != nil && root.Right.Val == root.Val {
        right += 1
        rightEqual = true
    }
    if !leftEqual && !rightEqual {
        // 都不相同,返回0
        return 0
    }else if(leftEqual && !rightEqual){
        return left
    }else if (!leftEqual && rightEqual){
        return right
    }
    return int(math.Max(float64(left),float64(right)))
}

func Leave2Leave(root *TreeNode) int {
    // 叶子到叶子节点
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    }
    r := 0
    // 两边的子树值和本身相等才算
    if root.Left != nil && root.Right != nil && root.Left.Val == root.Val && root.Right.Val == root.Val {
        // 取得两个子树的直达长度
        left := Root2Leaves(root.Left)
        right := Root2Leaves(root.Right)
        r += left + right + 2
    }else {
        r = int(math.Max(float64(Leave2Leave(root.Left)),float64(Leave2Leave(root.Right))))
    }
    return r
}
func longestUnivaluePath(root *TreeNode) int {
    // 两种情况,1.从根节点叶子的路径长度 2.叶子节点到叶子节点的路径长度
    root2LeaveLength := Root2Leaves(root)
    leave2LeaveLength := Leave2Leave(root)
    
    if root2LeaveLength > leave2LeaveLength{
        return root2LeaveLength
    }
    return leave2LeaveLength
}


func fib(n int) int {
    if n <= 1{
        return n
    }
    nums := make([]int64,n+1)
    for i:=0;i <= n;i++ {
        if i == 0 {
            nums[i] = 0
        }else if i == 1 {
            nums[i] = 1
        }else{
            nums[i] = nums[i-1]%1000000007 + nums[i-2]%1000000007
        }
    }
    
    return int(nums[n]%1000000007)
}
func move(n int, A *[]int, B *[]int, C *[]int){
    if n == 0 {
        return
    }
    if n == 1 {
        // 一个盘子,只需要从A->C
        *C = append(*C,(*A)[len(*A)-1])
        (*A) = (*A)[:len(*A)-1]
    }else{
        // 2以上个盘子 A->B A->C B->C
        move(n-1,A,C,B)
        move(1,A,B,C)
        move(n-1,B,A,C)
    }
}
func hanota(A []int, B []int, C []int) []int {
    if A == nil{
        return nil
    }
    move(len(A),&A,&B,&C)
    return C
}

func missingNumber(nums []int) int {
    first,last,mid := 0,len(nums)-1,(len(nums)/2)
    for first <= last {
        mid = (first + last)/2
        if mid == nums[mid] {
            first = mid + 1
        }else{
            last = mid -1
        }
    }
    return first
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }else if p != nil && q != nil {
        if p.Val == q.Val {
            return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
        }
    }
    return false
}
func isSymmetricTrees(p *TreeNode,q *TreeNode) bool{
    if p ==nil && q == nil {
        return true;
    }else if p!= nil && q != nil {
        if  p.Val == q.Val {
            return isSymmetricTrees(p.Left,q.Right) && isSymmetricTrees(p.Right,q.Left)
        }
    }
    return false
}
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSymmetricTrees(root.Left,root.Right)
}
func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left),maxDepth(root.Right))
}