package linkedList

import (
	"iter"
)

type NodeValues interface {
}

type Node[T NodeValues] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

func NewNode[T NodeValues](value T, prev *Node[T], next *Node[T]) *Node[T] {
	return &Node[T]{value, prev, next}
}

func (n *Node[T]) EachValues() iter.Seq[T] {
	return func(yield func(T) bool) {
		var nodeBase *Node[T] = n
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

func (n *Node[T]) EachNodes() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		var Node *Node[T] = n
		for {
			if Node == nil {
				break
			}
			stop := yield(Node)
			if !stop {
				break
			}
			Node = Node.Next
		}
	}
}
