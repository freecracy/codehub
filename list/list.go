package list

type listNode struct {
	prev, next *listNode
	val        int
}

// 两数相加,非递归版
func addTwoNumbers(list1 *listNode, list2 *listNode) *listNode {
	var a listNode

	b := new(listNode)
	carry := 0
	i := 0

	for {
		b.val = (list1.val + list2.val + carry) % 10
		carry = (list1.val + list2.val + carry) / 10

		if list1.next == nil && list2.next == nil && carry == 0 {
			a.val = b.val
			a.next = nil
			break
		}

		if list1 != nil {
			list1 = list1.next
		} else {
			list1.val = 0
			list1.next = nil
		}

		if list2 != nil {
			list2 = list2.next
		} else {
			list2.val = 0
			list2.next = nil
		}

		b.next = new(listNode)

		if i == 0 {
			a.next = b.next
		}

		b = b.next
		i++

	}
	return &a
}

// 递归两数相加
func addTwoNumbers2(list1, list2 *listNode, carry int) *listNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	b := new(listNode)
	b.val = (list1.val + list2.val + carry) % 10
	carry = (list1.val + list2.val + carry) / 10
	b.next = addTwoNumbers2(list1.next, list2.next, carry)
	return b
}
