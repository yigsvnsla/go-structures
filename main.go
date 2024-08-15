package main

import (
	pkg "go-structures/pkg"
)

func main() {
	players := [4]*pkg.Player{
		pkg.NewPlayer("p1", 1),
		pkg.NewPlayer("p2", 2),
		pkg.NewPlayer("p3", 3),
		pkg.NewPlayer("p4", 4),
	}

	table := pkg.CreateTable(
		players,
		4444,
	)

	pkg.PrintStruct(*table)
}
