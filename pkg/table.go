package pkg

import (
	"fmt"
	CardsEnum "go-structures/pkg/enums"
	"math/rand"
)

type Table struct {
	deck    [40]*Card
	cards   []*Card
	players *PlayerList
	id      uint
}

func CreateTable(playersTable *PlayerList, idTable uint) *Table {
	return &Table{
		cards:   make([]*Card, 0),
		players: playersTable,
		id:      idTable,
	}
}

func (t *Table) GenerateDeck() {
	var index uint8 = 0
	for _, vv := range CardsEnum.CardValuesMap() {
		for _, tv := range CardsEnum.CardTypeMap() {
			t.deck[index] = CreateCard(vv, tv)
			index++
		}
	}
}

func (t *Table) ShuffleDeck() {
	var lengthDeck int = len(t.deck)
	var cardTemp *Card = nil
	for i := lengthDeck; i > 1; i-- {
		indexRand := rand.Intn(i) + 1
		cardTemp = t.deck[indexRand]
		t.deck[indexRand] = t.deck[i-1]
		t.deck[i-1] = cardTemp
	}
}

func (t *Table) Show() {
	valuesIndex := CardsEnum.CardValuesMapIndex()
	typesIndex := CardsEnum.CardTypeMapIndex()
	for _, card := range t.deck {
		fmt.Println(valuesIndex[card.Value()], typesIndex[card.Type()])
	}
}
