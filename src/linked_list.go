package src

// Linked lists have one root node.
type LinkedList[T any] struct {
	RootNode *LinkedListNode[T]
}

// Linked list nodes contain the data and the pointer for the next node.
type LinkedListNode[T any] struct {
	Data     T
	NextNode *LinkedListNode[T]
}
