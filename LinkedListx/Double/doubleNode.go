package DoubleLinkedList

import (
	ll "go-structures/LinkedListx"
	"iter"
)

type DoubleNodeValues interface {
	ll.NodeBaseValues
}

type DoubleNode[T DoubleNodeValues] struct {
	ll.NodeBase[T]
	Next *DoubleNode[T]
	Prev *DoubleNode[T]
}

func NewDoubleNode[T DoubleNodeValues](value T, next *DoubleNode[T], prev *DoubleNode[T]) *DoubleNode[T] {
	node := &DoubleNode[T]{Next: next, Prev: prev, NodeBase: *ll.NewNodeBase(value, nil)}
	if prev != nil {
		prev.Next = node
	}

	if next != nil {
		next.Prev = node
	}

	if prev == nil && next == nil {
		node.Next = node
		node.Prev = node
	}

	return node
}

func (n *DoubleNode[T]) EachNodesRight() iter.Seq[*DoubleNode[T]] {
	return func(yield func(*DoubleNode[T]) bool) {
		start := n
		node := n
		for {
			// Si llegamos de nuevo al nodo inicial, detener la iteración
			if node == nil || node == start && node != n {
				break
			}
			// Llamar a yield con el nodo actual
			stop := yield(node)
			if !stop {
				break
			}
			// Mover al siguiente nodo
			node = node.Prev
			// Si la lista es circular, evitar el bucle infinito
			if node == start {
				break
			}
		}
	}
}

func (n *DoubleNode[T]) EachNodes() iter.Seq[*DoubleNode[T]] {
	return func(yield func(*DoubleNode[T]) bool) {
		start := n
		node := n
		for {
			// Si llegamos de nuevo al nodo inicial, detener la iteración
			if node == nil || node == start && node != n {
				break
			}
			// Llamar a yield con el nodo actual
			stop := yield(node)
			if !stop {
				break
			}
			// Mover al siguiente nodo
			node = node.Next
			// Si la lista es circular, evitar el bucle infinito
			if node == start {
				break
			}
		}
	}
}
