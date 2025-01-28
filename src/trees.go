package src

type Tree[T any] struct {
	RootNode TreeNode[T]
}

type TreeNode[T any] struct {
	Data     T
	Children []TreeNode[T]
}
