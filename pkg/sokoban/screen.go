// SPDX-License-Identifier: CC-BY-SA-4.0
package sokoban

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ooxi/sokoban-ebiten-assets/pkg/assets"
)

type Screen struct {
	image  *ebiten.Image
	assets *assets.SokobanAssets
}

func (screen Screen) at(tileX, tileY int) *ebiten.DrawImageOptions {
	x := float64(tileX * screen.assets.TileWidth)
	y := float64(tileY * screen.assets.TileHeight)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	return op
}

func (screen Screen) DrawCharacter1(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter1, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter2(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter2, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter3(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter3, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter4(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter4, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter5(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter5, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter6(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter6, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter7(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter7, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter8(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter8, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter9(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter9, screen.at(tileX, tileY))
}

func (screen Screen) DrawCharacter10(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCharacter10, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateBeige(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateBeige, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateBlack(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateBlack, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateBlue(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateBlue, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateBrown(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateBrown, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateGray(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateGray, screen.at(tileX, tileY))
}

func (screen Screen) DrawCratePurple(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCratePurple, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateRed(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateRed, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateYellow(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateYellow, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkBeige(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkBeige, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkBlack(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkBlack, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkBlue(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkBlue, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkBrown(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkBrown, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkGray(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkGray, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkPurple(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkPurple, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkRed(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkRed, screen.at(tileX, tileY))
}

func (screen Screen) DrawCrateDarkYellow(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageCrateDarkYellow, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointBeige(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointBeige, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointBlack(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointBlack, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointBlue(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointBlue, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointBrown(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointBrown, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointGray(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointGray, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointPurple(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointPurple, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointRed(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointRed, screen.at(tileX, tileY))
}

func (screen Screen) DrawEndpointYellow(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageEndpointYellow, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundConcrete(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundConcrete, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundDirt(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundDirt, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundGrass(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundGrass, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundSand(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundSand, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundGravelConcrete(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundGravelConcrete, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundGravelDirt(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundGravelDirt, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundGravelGrass(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundGravelGrass, screen.at(tileX, tileY))
}

func (screen Screen) DrawGroundGravelSand(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageGroundGravelSand, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallBeige(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallBeige, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallBlack(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallBlack, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallBrown(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallBrown, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallGray(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallGray, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallRoundBeige(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallRoundBeige, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallRoundBlack(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallRoundBlack, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallRoundBrown(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallRoundBrown, screen.at(tileX, tileY))
}

func (screen Screen) DrawWallRoundGray(tileX, tileY int) {
	screen.image.DrawImage(screen.assets.ImageWallRoundGray, screen.at(tileX, tileY))
}
