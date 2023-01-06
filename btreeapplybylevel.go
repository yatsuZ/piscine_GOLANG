package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		if root.Parent == nil {
			f(root.Data)
		}
		if root.Left != nil {
			f(root.Left.Data)
			BTreeApplyByLevel(root.Right, f)
		}
		if root.Right != nil {
			f(root.Right.Data)
			BTreeApplyByLevel(root.Left, f)
		}
	}
}
