package link_list

import (
	ll "go-structures/link-list"
)

type SimpleNodeValues interface {
	ll.NodeBaseValues
}

type SimpleNode[T SimpleNodeValues] struct {
	ll.NodeBase[T]
	Next *SimpleNode[T]
}

func SimpleNewNode[T SimpleNodeValues](value T, next *SimpleNode[T]) *SimpleNode[T] {
	node := &SimpleNode[T]{NodeBase: *ll.NewNodeBase[T](value, nil), Next: next}
	return node
}
