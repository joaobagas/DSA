package dstr

// Linked list nodes contain the data and the pointer for the next node.
type LinkedListNode[T any] struct {
	data     T
	nextNode *LinkedListNode[T]
}

func GetData[T any](node LinkedListNode[T]) T {
	return node.data
}

func NextNode[T any](node *LinkedListNode[T]) LinkedListNode[T] {
	return *node.nextNode
}
