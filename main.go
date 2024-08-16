package main

import (
	pkg "go-structures/pkg"
	"go-structures/pkg/utils"
)

func main() {
	players := []pkg.Player{
		*pkg.NewPlayer("p1", 1),
		*pkg.NewPlayer("p2", 2),
		*pkg.NewPlayer("p3", 3),
		*pkg.NewPlayer("p4", 4),
	}

	playerList := pkg.NewPlayerList()
	playerList.Append(players...)

	table := pkg.CreateTable(playerList, 4444)
	table.GenerateDeck()
	table.ShuffleDeck()
	// table.Show()
	utils.PrintStruct(*table)
}
