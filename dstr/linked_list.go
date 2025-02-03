package dstr

// Linked list nodes contain the data and the pointer for the next node.
type LinkedListNode[T any] struct {
	data     T
	nextNode *LinkedListNode[T]
}

func CreateLinkedList[T any](newData T, newNextNode *LinkedListNode[T]) LinkedListNode[T] {
	var head LinkedListNode[T]

	if newNextNode != nil {
		head = LinkedListNode[T]{
			data:     newData,
			nextNode: newNextNode,
		}
	} else {
		head = LinkedListNode[T]{
			data:     newData,
			nextNode: nil,
		}
	}

	return head
}

// Gets the data inside the current node.
func GetData[T any](node LinkedListNode[T]) T {
	return node.data
}

// Gets the nect node in the linked list.
func NextNode[T any](node LinkedListNode[T]) LinkedListNode[T] {
	return *node.nextNode
}

func AddNode[T any](node *LinkedListNode[T], nodeToAdd *LinkedListNode[T]) {
	var curNode LinkedListNode[T]
	for {
		if curNode.nextNode != nil {
			curNode = *curNode.nextNode
		} else {
			break
		}
	}
	curNode.nextNode = nodeToAdd
}

func IsNextNode[T any](node LinkedListNode[T]) bool {
	return node.nextNode == nil
}
