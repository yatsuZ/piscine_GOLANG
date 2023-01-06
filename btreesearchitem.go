package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root != nil {
		if root.Data > elem {
			return BTreeSearchItem(root.Left, elem)
		} else if root.Data == elem {
			return root
		} else {
			return BTreeSearchItem(root.Right, elem)
		}
	}
	return nil
}
