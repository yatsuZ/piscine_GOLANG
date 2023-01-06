package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root != nil {
		if root.Left != nil && root.Data <= root.Left.Data {
			return false
		} else if root.Right != nil && root.Data > root.Right.Data {
			return false
		}
		if BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right) {
			return true
		} else {
			return false
		}
	}
	return true
}
