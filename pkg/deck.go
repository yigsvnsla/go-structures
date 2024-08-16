package pkg

// import (
// 	"fmt"
// 	CardsEnum "go-structures/pkg/enums"
// 	"math/rand"
// )

// type Deck struct {
// 	cards
// }

// func GenerateDeck() *Deck {
// 	d := new(Deck)
// 	var index uint8 = 0
// 	for _, vv := range CardsEnum.CardValuesMap() {
// 		for _, tv := range CardsEnum.CardTypeMap() {
// 			d.cards[index] = CreateCard(vv, tv)
// 			index++
// 		}
// 	}
// 	return d
// }

// func (d *Deck) Shuffle() {
// 	var lengthDeck int = len(d.cards)
// 	var cardTemp *Card = nil
// 	for i := lengthDeck; i > 1; i-- {
// 		indexRand := rand.Intn(10) + 1
// 		cardTemp = d.cards[indexRand]
// 		d.cards[indexRand] = d.cards[i-1]
// 		d.cards[i-1] = cardTemp
// 	}
// }

// func (d *Deck) Show() {
// 	valuesIndex := CardsEnum.CardValuesMapIndex()
// 	typesIndex := CardsEnum.CardTypeMapIndex()
// 	for _, card := range d.cards {
// 		fmt.Println(valuesIndex[card.Value()], typesIndex[card.Type()])
// 	}
// }
