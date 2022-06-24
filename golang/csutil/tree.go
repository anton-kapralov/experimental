package csutil

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tn TreeNode) String() string {
	return strconv.Itoa(tn.Val)
}

func InOrder(tn *TreeNode, consume func(*TreeNode)) {
	var d Deque
	stackLefts(tn, &d)
	for !d.IsEmpty() {
		tn = d.RemoveLast().(*TreeNode)
		consume(tn)
		stackLefts(tn.Right, &d)
	}
}

func stackLefts(tn *TreeNode, d *Deque) {
	for tn != nil {
		d.AddLast(tn)
		tn = tn.Left
	}
}
