package SimpleLinkedList

import (
	"fmt"
	b "go-structures/book"
	"testing"
)

type MyTestIntValue struct {
	int
}

func (T MyTestIntValue) GetValue() string {
	return fmt.Sprint(T.int)
}

func TestLinkListAppend(t *testing.T) {
	// Crear una lista enlazada vacía
	linkList := SimpleLinkedList[b.Book]{}

	// Verificar que la lista esté vacía
	if linkList.GetSize() != 0 {
		t.Errorf("Esperaba que la lista estuviera vacía, pero tiene %d elementos", linkList.GetSize())
	}

	// Agregar elementos a la lista enlazada
	linkList.Append(b.NewBook(1, "book 1"))
	linkList.Append(b.NewBook(2, "book 2"))
	linkList.Append(b.NewBook(3, "book 3"))
	linkList.Append(b.NewBook(6, "book 6"))

	// Verificar que la lista tenga el número correcto de elementos
	expectedLen := 4
	if linkList.GetSize() == 0 {
		t.Errorf("Esperaba que la lista tuviera %d elementos, pero tiene %d", expectedLen, linkList.GetSize())
	}

	// Verificar que los elementos sean correctos
	expectedValues := []b.Book{
		b.NewBook(1, "book 1"),
		b.NewBook(2, "book 2"),
		b.NewBook(3, "book 3"),
		b.NewBook(6, "book 6"),
	}

	current := linkList.GetHead()
	for i, expectedBook := range expectedValues {
		if current == nil {
			t.Fatalf("Esperaba que el nodo %d contuviera %v, pero era nil", i, expectedBook)
		}
		if current.Value != expectedBook {
			t.Errorf("Esperaba que el nodo %d contuviera %v, pero obtuvo %v", i, expectedBook, current.Value)
		}
		current = current.Next
	}

	// Verificar que no haya más elementos
	if current != nil {
		t.Errorf("Esperaba que no hubiera más elementos en la lista, pero hay más.")
	}
}

// Append single value to an empty list
func TestAppendSingleValueToEmptyList(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	node, err := ll.Append(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.head == nil || ll.tail == nil {
		t.Fatalf("head or tail is nil")
	}
	if ll.head.Value != 1 || ll.tail.Value != 1 {
		t.Fatalf("expected head and tail value to be 1, got head: %v, tail: %v", ll.head.Value, ll.tail.Value)
	}
	if node.Value != 1 {
		t.Fatalf("expected node value to be 1, got %v", node.Value)
	}
}

// Append multiple values to an empty list
func TestAppendMultipleValuesToEmptyList(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	_, err := ll.Append(1, 2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.head == nil || ll.tail == nil {
		t.Fatalf("head or tail is nil")
	}
	if ll.head.Value != 1 || ll.tail.Value != 3 {
		t.Fatalf("expected head value to be 1 and tail value to be 3, got head: %v, tail: %v", ll.head.Value, ll.tail.Value)
	}
}

// Append single value to a non-empty list
func TestAppendSingleValueToNonEmptyList(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	ll.Append(1)
	node, err := ll.Append(2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.tail.Value != 2 {
		t.Fatalf("expected tail value to be 2, got %v", ll.tail.Value)
	}
	if node.Value != 2 {
		t.Fatalf("expected node value to be 2, got %v", node.Value)
	}
}

// Append multiple values to a non-empty list
func TestAppendMultipleValuesToNonEmptyList(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	ll.Append(1)
	_, err := ll.Append(2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.tail.Value != 3 {
		t.Fatalf("expected tail value to be 3, got %v", ll.tail.Value)
	}
}

// Verify the list size after appending values
func TestVerifyListSizeAfterAppendingValues(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	ll.Append(1, 2, 3)
	if ll.size != 3 {
		t.Fatalf("expected list size to be 3, got %v", ll.size)
	}
}

// Append no values to the list
func TestAppendNoValuesToList(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	_, err := ll.Append()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.size != 0 {
		t.Fatalf("expected list size to be 0, got %v", ll.size)
	}
}

// Append a nil value to the list
func TestAppendNilValueToList(t *testing.T) {
	ll := &SimpleLinkedList[*int]{}
	var val *int = nil
	_, err := ll.Append(val)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.size != 1 {
		t.Fatalf("expected list size to be 1, got %v", ll.size)
	}
}

// Append values to a list with existing nil nodes
func TestAppendValuesToListWithExistingNilNodes(t *testing.T) {
	ll := &SimpleLinkedList[*int]{}
	var val *int = nil
	ll.Append(val)
	_, err := ll.Append(new(int))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.size != 2 {
		t.Fatalf("expected list size to be 2, got %v", ll.size)
	}
}

// Append values to a list with only one node
func TestAppendValuesToListWithOneNode(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	ll.Append(1)
	_, err := ll.Append(2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.size != 3 {
		t.Fatalf("expected list size to be 3, got %v", ll.size)
	}
}

// Append values to a list with maximum capacity
func TestAppendValuesToListWithMaxCapacity(t *testing.T) {
	const maxCapacity = 1000
	ll := &SimpleLinkedList[int]{}
	for i := 0; i < maxCapacity; i++ {
		_, err := ll.Append(i)
		if err != nil {
			t.Fatalf("unexpected error at iteration %d: %v", i, err)
		}
	}
	if ll.size != maxCapacity {
		t.Fatalf("expected list size to be %d, got %v", maxCapacity, ll.size)
	}
}

// Ensure the tail pointer is updated correctly after appending
func TestTailPointerUpdatedCorrectlyAfterAppending(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	_, err := ll.Append(1, 2, 3)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ll.tail.Value != 3 {
		t.Fatalf("expected tail value to be 3, got %v", ll.tail.Value)
	}
}

// Verify the head pointer remains unchanged when appending to a non-empty list
func TestHeadPointerUnchangedWhenAppendingToNonEmptyList(t *testing.T) {
	ll := &SimpleLinkedList[int]{}
	_, _ = ll.Append(1)
	headBefore := ll.head
	_, _ = ll.Append(2, 3)

	if headBefore != ll.head {
		t.Fatalf("expected head pointer to remain unchanged")
	}
}
