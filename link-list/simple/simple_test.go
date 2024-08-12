package link_list

import (
	b "go-structures/book"
	"testing"
)

func TestLinkListAppend(t *testing.T) {
	// Crear una lista enlazada vacía
	linkList := SimpleLinkList[b.Book]{}

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
