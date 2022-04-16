// SPDX-License-Identifier: CC-BY-SA-4.0
package sokoban

type Game interface {

	// Update updates a game by one tick.
	Update() error

	// Draw draw the game screen. The given argument represents a screen image.
	Draw(screen Screen)

	// Return number of tiles to be displayed at once in X and Y
	Layout() (tilesX, tilesY int)
}
