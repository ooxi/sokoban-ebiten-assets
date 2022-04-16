// SPDX-License-Identifier: CC-BY-SA-4.0
package main

import "github.com/ooxi/sokoban-ebiten-assets/pkg/sokoban"

type Game struct {
}

func (game *Game) Update() error {
	return nil
}

func (game *Game) Draw(screen sokoban.Screen) {
	screen.DrawGroundGravelConcrete(0, 0)
	screen.DrawCharacter1(2, 3)
	screen.DrawGroundSand(4, 6)
}

func (game *Game) Layout() (int, int) {
	return 5, 7
}

func main() {
	game := &Game{}
	sokoban.RunGame("Simple Demo", game)
}
