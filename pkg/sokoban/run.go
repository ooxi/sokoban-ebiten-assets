// SPDX-License-Identifier: CC-BY-SA-4.0
package sokoban

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ooxi/sokoban-ebiten-assets/pkg/assets"
)

type gameContainer struct {
	game   Game
	assets *assets.SokobanAssets
}

func (container *gameContainer) Update() error {
	return container.game.Update()
}

func (container *gameContainer) Draw(screen *ebiten.Image) {
	container.game.Draw(Screen{screen, container.assets})
}

func (container *gameContainer) layout() (screenWidth, screenHeight int) {
	x, y := container.game.Layout()
	return x * container.assets.TileWidth, y * container.assets.TileHeight
}

func (container *gameContainer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return container.layout()
}

func RunGame(title string, game Game) error {
	assets, err := assets.NewSokobanAssets()

	if nil != err {
		return fmt.Errorf("failed loading sokoban assets: %w", err)
	}

	container := &gameContainer{game, assets}

	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(container.layout())

	return ebiten.RunGame(container)
}
