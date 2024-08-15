package pkg

import (
	CardsEnum "go-structures/pkg/enums"
)

type Card struct {
	v CardsEnum.CardValue
	t CardsEnum.CardType
}

func CreateCard(cardValue CardsEnum.CardValue, cardForm CardsEnum.CardType) *Card {
	return &Card{v: cardValue, t: cardForm}
}

func (c *Card) Value() CardsEnum.CardValue {
	return c.v
}

func (c *Card) Type() CardsEnum.CardType {
	return c.t
}
