package dstr

// A tree node contains data and other three nodes.
type TreeNode[T any] struct {
	Data     T
	Children []*TreeNode[T]
}
