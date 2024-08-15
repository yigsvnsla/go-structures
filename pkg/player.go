package pkg

type Player struct {
	cards [3]*Card
	name  string
	id    uint
}

func NewPlayer(namePlayer string, idPlayer uint) *Player {
	return &Player{name: namePlayer, id: idPlayer}
}
