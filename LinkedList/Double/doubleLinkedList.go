package DoubleLinkedList

import "fmt"

type DoubleLinkedListValues interface {
}

type DoubleLinkedList[T DoubleLinkedListValues] struct {
	Head *DoubleNode[T]
	Tail *DoubleNode[T]
	List *DoubleNode[T]
	Size int
}

func (ll *DoubleLinkedList[T]) Append(values ...T) (*DoubleNode[T], error) {
	var stageNodes *DoubleNode[T]

	for v := range values {
		// newNode := NewDoubleNode()
		fmt.Println(v)
	}

	return stageNodes, nil
}
