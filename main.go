package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// IsBalanced reports whether every node's subtrees differ in height by at
// most one. The height recursion returns -1 as an "already unbalanced"
// signal so the walk short-circuits instead of re-measuring.
func IsBalanced(root *TreeNode) bool {
	return checkedHeight(root) >= 0
}

func checkedHeight(n *TreeNode) int {
	if n == nil {
		return 0
	}
	l := checkedHeight(n.Left)
	if l < 0 {
		return -1
	}
	r := checkedHeight(n.Right)
	if r < 0 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	balanced := tn(3, tn(9, nil, nil), tn(20, tn(15, nil, nil), tn(7, nil, nil)))
	skewed := tn(1, tn(2, tn(3, nil, nil), nil), nil)
	fmt.Println(IsBalanced(balanced), IsBalanced(skewed))
}
