package LinkedList

import (
	"iter"
)

type NodeBaseValues interface {
}

type NodeBase[T NodeBaseValues] struct {
	Value T
	Next  *NodeBase[T]
}

func NewNodeBase[T NodeBaseValues](value T, next *NodeBase[T]) *NodeBase[T] {
	return &NodeBase[T]{value, next}
}

func (n *NodeBase[T]) EachValues() iter.Seq[T] {
	return func(yield func(T) bool) {
		var nodeBase *NodeBase[T] = n
		for {
			if nodeBase == nil {
				break
			}
			stop := yield(nodeBase.Value)
			if !stop {
				break
			}
			nodeBase = nodeBase.Next
		}
	}
}

func (n *NodeBase[T]) EachNodes() iter.Seq[*NodeBase[T]] {
	return func(yield func(*NodeBase[T]) bool) {
		var nodeBase *NodeBase[T] = n
		for {
			if nodeBase == nil {
				break
			}
			stop := yield(nodeBase)
			if !stop {
				break
			}
			nodeBase = nodeBase.Next
		}
	}
}
