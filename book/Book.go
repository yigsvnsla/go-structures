package Book

import "fmt"

type Book struct {
	id   int
	name string
}

func (b Book) GetValue() string {
	return fmt.Sprintf("%d ", b.id)
}

func (b *Book) SetId(id int) {
	b.id = id
}

func (b *Book) GetId() int {
	return b.id
}

func (b *Book) SetName(name string) {
	b.name = name
}

func (b *Book) GetName() string {
	return b.name
}

func NewBook(id int, name string) Book {
	return Book{id, name}
}
