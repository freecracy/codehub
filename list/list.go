package list

// ListNode ....
type ListNode struct {
	Prex, Next *ListNode
	Val        int
}

// 两数相加,非递归版
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := new(ListNode)
	p, q, curr := l1, l2, dummyHead
	carry := 0
	var x, y int
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}
	return dummyHead.Next
}

// 递归两数相加
func addTwoNumbers2(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	x, y := 0, 0
	if l1 != nil {
		x = l1.Val
	} else {
		l1 = &ListNode{Val: 0, Next: nil}
	}
	if l2 != nil {
		y = l2.Val
	} else {
		l2 = &ListNode{Val: 0, Next: nil}
	}
	sum := x + y
	nextNode := addTwoNumbers(l1.Next, l2.Next)
	if sum < 10 {
		return &ListNode{Val: sum, Next: nextNode}
	}
	tmpNode := &ListNode{Val: 1, Next: nil}
	return &ListNode{Val: sum - 10, Next: addTwoNumbers(nextNode, tmpNode)}
}
