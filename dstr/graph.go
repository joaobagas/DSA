package dstr

// A graph does not contain one root node.
type Graph[T any] struct {
	Graphs []GraphNode[T]
}

// Graph nodes can point to multiple Graph nodes and can be pointed at by multiple Graph nodes.
type GraphNode[T any] struct {
	Data     T
	Children []*GraphNode[T]
}
