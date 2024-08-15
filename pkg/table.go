package pkg

type Table struct {
	deck    *Deck
	cards   []*Card
	players [4]*Player
	id      uint
}

func CreateTable(playersTable [4]*Player, idTable uint) *Table {
	return &Table{
		deck:    GenerateDeck(),
		cards:   make([]*Card, 0),
		players: playersTable,
		id:      idTable,
	}
}
