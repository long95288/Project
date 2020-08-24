package leetcode

import (
    "fmt"
    "math"
    "sort"
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