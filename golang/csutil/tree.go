package csutil

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root TreeNode) String() string {
	var sb strings.Builder
	sb.WriteString("{")
	InOrder(&root, func(tn *TreeNode) {
		if sb.Len() > 1 {
			sb.WriteString(", ")
		}
		sb.WriteString(strconv.Itoa(tn.Val))
	})
	sb.WriteString("}")
	return sb.String()
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
