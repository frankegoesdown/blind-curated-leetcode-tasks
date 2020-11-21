package main

func main()

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	nextLevel := []*TreeNode{root}

	for len(nextLevel) > 0 {
		level := []*TreeNode{}
		current := []int{}
		for _, node := range nextLevel {
			current = append(current, node.Val)

			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		result = append(result, current)
		nextLevel = level
	}
	return result
}
