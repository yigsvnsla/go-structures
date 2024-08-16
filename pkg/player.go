package pkg

import (
	"go-structures/pkg/linkedList"
)

// type DeckPlayer struct {
// 	Deck
// 	// cards []*Card
// }

type Player struct {
	cards [3]*Card
	// deck  DeckPlayer
	name string
	id   uint
}

func NewPlayer(namePlayer string, idPlayer uint) *Player {
	return &Player{name: namePlayer, id: idPlayer}
}

type PlayerList struct {
	linkedList.Type[Player]
}

func NewPlayerList(players ...Player) *PlayerList {
	playerlist := new(PlayerList)
	playerlist.Append(players...)
	return playerlist
}

func (pl *PlayerList) Append(players ...Player) {
	for _, value := range players {
		pl.Type.Append(linkedList.NewNode(value, nil, nil))
	}
}
