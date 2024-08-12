package main

import (
	"fmt"
	b "go-structures/book"
	link_list "go-structures/link-list"
)

func main() {

	// newDoubleNode := doublelinklist.NewDoubleNode[b.Book](b.NewBook(1, "book 1"), nil, nil)

	// fmt.Println(newDoubleNode.NodeBase.Value.GetName())

	// // for val := range newDoubleNode.EachNodesRight() {
	// // 	fmt.Println(val.Value)
	// // }

	// newDoubleNode2 := doublelinklist.NewDoubleNode[b.Book](b.NewBook(2, "book 2"), newDoubleNode, newDoubleNode)

	// _ = doublelinklist.NewDoubleNode[b.Book](b.NewBook(3, "book 3"), newDoubleNode, newDoubleNode2)

	// for val := range newDoubleNode.EachNodesRight() {
	// 	fmt.Println(val.Value)
	// }

	// for val := range newDoubleNode3.EachNodes() {
	// 	fmt.Println(val.Value)
	// }

	// fmt.Println(NewDoubleNode.Value)

	// books1 := []b.Book{
	// 	b.NewBook(1, "book 1"),
	// 	b.NewBook(2, "book 2"),
	// 	b.NewBook(3, "book 3"),
	// }

	// books2 := []b.Book{
	// 	b.NewBook(4, "book 4"),
	// 	b.NewBook(5, "book 5"),
	// 	b.NewBook(6, "book 6"),
	// }

	// linkList := ll.Linklist[b.Book]{}

	// linkList.Append(books1...)
	// linkList.Append(books2...)

	// node, err := linkList.Find(func(b b.Book) bool {
	// 	return b.GetId() == 2
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(node)

	// deleteNode, err := linkList.Delete(func(b b.Book) bool {
	// 	return b.GetId() == 6
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("deleted ", deleteNode)

	// deleteNode, err := linkList.InsertBefore(func(b b.Book) bool {
	// 	return b.GetId() == 1
	// }, b.NewBook(10, "book 10"))

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("insert parent ", deleteNode)

	// linkList.Show()

	myBook := b.NewBook(10, "Harry potter")

	BaseNode := link_list.NewNodeBase(myBook, nil)

	LoR := b.NewBook(10, "Lord of the rings")

	BaseNode2 := link_list.NewNodeBase(LoR, nil)

	BaseNode.Next = BaseNode2

	fmt.Println(BaseNode)
}
