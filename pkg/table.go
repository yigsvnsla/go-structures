package pkg

import "go-structures/pkg/linkedList"

type Table struct {
	deck    *Deck
	cards   []*Card
	players linkedList.TypeI[Player]
	id      uint
}

func CreateTable(playersTable linkedList.TypeI[Player], idTable uint) *Table {
	return &Table{
		deck:    GenerateDeck(),
		cards:   make([]*Card, 0),
		players: playersTable,
		id:      idTable,
	}
}
