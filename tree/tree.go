package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 反转二叉树,递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Left = right
	root.Right = left
	return root
}

// 反转二叉树,迭代
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		temp := current.Left
		current.Left = current.Right
		current.Right = temp
		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return root
}
