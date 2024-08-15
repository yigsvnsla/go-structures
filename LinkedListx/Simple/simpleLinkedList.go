package SimpleLinkedList

import (
	"errors"
	"fmt"
)

const (
	NodeNotFound = "Node Not Found"
)

type SimpleLinkedListValues interface {
}

type SimpleLinkedList[T SimpleLinkedListValues] struct {
	head *SimpleNode[T]
	tail *SimpleNode[T]
	size int
}

func (ll *SimpleLinkedList[T]) GetSize() int {
	return ll.size
}

func (ll *SimpleLinkedList[T]) GetHead() *SimpleNode[T] {
	return ll.head
}

func (ll *SimpleLinkedList[T]) GetTail() *SimpleNode[T] {
	return ll.tail
}

func (ll *SimpleLinkedList[T]) Append(values ...T) (*SimpleNode[T], error) {
	var stageNodes *SimpleNode[T]
	for _, value := range values {
		newNode := SimpleNewNode(value, nil)
		if stageNodes == nil {
			stageNodes = newNode
		}
		if ll.head == nil {
			ll.head = newNode
			ll.tail = ll.head
		} else if ll.tail != nil {
			ll.tail.Next = newNode
			ll.tail = ll.tail.Next
		}
		ll.size++
	}
	return stageNodes, nil
}

func (ll *SimpleLinkedList[T]) Contains(predicate func(T) bool) bool {
	_, err := ll.Find(predicate)
	return err == nil
}

func (ll *SimpleLinkedList[T]) Delete(predicate func(T) bool) (*SimpleNode[T], error) {
	currentNode := ll.head
	if currentNode == nil {
		return nil, errors.New("Head Null")
	}
	var preventNode *SimpleNode[T]
	for value := range currentNode.EachNodes() {
		node := &SimpleNode[T]{NodeBase: *value}
		if predicate(node.Value) {
			if preventNode == nil {
				ll.head = node.Next
			} else {
				preventNode.Next = node.Next
			}
			return node, nil
		}
		preventNode = node
	}
	return nil, errors.New("Node Not Found")
}

func (ll *SimpleLinkedList[T]) Find(predicate func(T) bool) (*SimpleNode[T], error) {
	currentNode := ll.head
	if currentNode == nil {
		return nil, errors.New("Head Null")
	}
	for value := range currentNode.EachNodes() {
		if predicate(value.Value) {
			return &SimpleNode[T]{NodeBase: *value}, nil
		}
	}
	return nil, errors.New("Node Not Found")
}

func (ll *SimpleLinkedList[T]) FindParent(predicate func(T) bool) (*SimpleNode[T], error) {
	currentNode := ll.head
	if currentNode == nil {
		return nil, errors.New("Head Null")
	}
	var preventNode *SimpleNode[T]
	for value := range currentNode.EachNodes() {
		if predicate(value.Value) {
			if preventNode == nil {
				return nil, errors.New("The head has not father")
			}
			return preventNode, nil
		}
		preventNode = &SimpleNode[T]{NodeBase: *value}
	}
	return nil, errors.New(NodeNotFound)
}

func (ll *SimpleLinkedList[T]) InsertBefore(predicate func(T) bool, newValue T) (*SimpleNode[T], error) {
	headNode := ll.head
	parentNode, err := ll.FindParent(predicate)
	if err != nil {
		if err.Error() == NodeNotFound {
			return nil, err
		}
	}
	if err != nil && parentNode == nil {
		ll.head = SimpleNewNode(newValue, headNode)
		return ll.head, nil
	}
	parentNode.Next = SimpleNewNode(newValue, parentNode.Next)
	return parentNode.Next, nil
}

func (ll *SimpleLinkedList[T]) InsertAfter(predicate func(T) bool, newValue T) (*SimpleNode[T], error) {
	node, err := ll.Find(predicate)
	if err != nil {
		fmt.Println("error insert after")
	}
	newNode := SimpleNewNode(newValue, node)
	node.Next = newNode
	if node == ll.tail {
		ll.tail = newNode
	}
	return newNode, err
}

func (ll *SimpleLinkedList[T]) Replace(predicate func(T) bool, newValue T) (*SimpleNode[T], error) {
	node, err := ll.Find(predicate)
	if err != nil {
		return nil, err
	}
	node.Value = newValue
	return node, nil
}

func (ll *SimpleLinkedList[T]) Show() {
	for v := range ll.head.EachValues() {
		fmt.Print("-> ", v)
	}
	fmt.Println()
}
