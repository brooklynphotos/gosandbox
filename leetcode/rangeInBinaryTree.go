// https://leetcode.com/problems/range-sum-of-bst/

package leetcode

func rangeSumBST(root *TreeNode, L int, R int) int {
	return rangeSumSubproblem(L, R, 0, []*TreeNode{root})
}

func rangeSumSubproblem(L int, R int, sum int, queue []*TreeNode) int {
	if len(queue) == 0 {
		return sum
	}
	// pop out the first item
	root := queue[0]
	queue = queue[1:]
	switch {
	case root.Val < L:
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	case root.Val > R:
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
	default:
		sum += root.Val
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
	}
	return rangeSumSubproblem(L, R, sum, queue)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
