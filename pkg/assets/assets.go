// SPDX-License-Identifier: CC-BY-SA-4.0
package assets

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type SokobanAssets struct {
	TileWidth  int
	TileHeight int

	ImageCharacter1  *ebiten.Image
	ImageCharacter2  *ebiten.Image
	ImageCharacter3  *ebiten.Image
	ImageCharacter4  *ebiten.Image
	ImageCharacter5  *ebiten.Image
	ImageCharacter6  *ebiten.Image
	ImageCharacter7  *ebiten.Image
	ImageCharacter8  *ebiten.Image
	ImageCharacter9  *ebiten.Image
	ImageCharacter10 *ebiten.Image

	ImageCrateBeige  *ebiten.Image
	ImageCrateBlack  *ebiten.Image
	ImageCrateBlue   *ebiten.Image
	ImageCrateBrown  *ebiten.Image
	ImageCrateGray   *ebiten.Image
	ImageCratePurple *ebiten.Image
	ImageCrateRed    *ebiten.Image
	ImageCrateYellow *ebiten.Image

	ImageCrateDarkBeige  *ebiten.Image
	ImageCrateDarkBlack  *ebiten.Image
	ImageCrateDarkBlue   *ebiten.Image
	ImageCrateDarkBrown  *ebiten.Image
	ImageCrateDarkGray   *ebiten.Image
	ImageCrateDarkPurple *ebiten.Image
	ImageCrateDarkRed    *ebiten.Image
	ImageCrateDarkYellow *ebiten.Image

	ImageEndpointBeige  *ebiten.Image
	ImageEndpointBlack  *ebiten.Image
	ImageEndpointBlue   *ebiten.Image
	ImageEndpointBrown  *ebiten.Image
	ImageEndpointGray   *ebiten.Image
	ImageEndpointPurple *ebiten.Image
	ImageEndpointRed    *ebiten.Image
	ImageEndpointYellow *ebiten.Image

	ImageGroundConcrete *ebiten.Image
	ImageGroundDirt     *ebiten.Image
	ImageGroundGrass    *ebiten.Image
	ImageGroundSand     *ebiten.Image

	ImageGroundGravelConcrete *ebiten.Image
	ImageGroundGravelDirt     *ebiten.Image
	ImageGroundGravelGrass    *ebiten.Image
	ImageGroundGravelSand     *ebiten.Image

	ImageWallBeige *ebiten.Image
	ImageWallBlack *ebiten.Image
	ImageWallBrown *ebiten.Image
	ImageWallGray  *ebiten.Image

	ImageWallRoundBeige *ebiten.Image
	ImageWallRoundBlack *ebiten.Image
	ImageWallRoundBrown *ebiten.Image
	ImageWallRoundGray  *ebiten.Image
}

var (
	//go:embed assets/Character1.png
	resourceCharacter1 []byte

	//go:embed assets/Character2.png
	resourceCharacter2 []byte

	//go:embed assets/Character3.png
	resourceCharacter3 []byte

	//go:embed assets/Character4.png
	resourceCharacter4 []byte

	//go:embed assets/Character5.png
	resourceCharacter5 []byte

	//go:embed assets/Character6.png
	resourceCharacter6 []byte

	//go:embed assets/Character7.png
	resourceCharacter7 []byte

	//go:embed assets/Character8.png
	resourceCharacter8 []byte

	//go:embed assets/Character9.png
	resourceCharacter9 []byte

	//go:embed assets/Character10.png
	resourceCharacter10 []byte

	//go:embed assets/Crate_Beige.png
	resourceCrateBeige []byte

	//go:embed assets/Crate_Black.png
	resourceCrateBlack []byte

	//go:embed assets/Crate_Blue.png
	resourceCrateBlue []byte

	//go:embed assets/Crate_Brown.png
	resourceCrateBrown []byte

	//go:embed assets/Crate_Gray.png
	resourceCrateGray []byte

	//go:embed assets/Crate_Purple.png
	resourceCratePurple []byte

	//go:embed assets/Crate_Red.png
	resourceCrateRed []byte

	//go:embed assets/Crate_Yellow.png
	resourceCrateYellow []byte

	//go:embed assets/CrateDark_Beige.png
	resourceCrateDarkBeige []byte

	//go:embed assets/CrateDark_Black.png
	resourceCrateDarkBlack []byte

	//go:embed assets/CrateDark_Blue.png
	resourceCrateDarkBlue []byte

	//go:embed assets/CrateDark_Brown.png
	resourceCrateDarkBrown []byte

	//go:embed assets/CrateDark_Gray.png
	resourceCrateDarkGray []byte

	//go:embed assets/CrateDark_Purple.png
	resourceCrateDarkPurple []byte

	//go:embed assets/CrateDark_Red.png
	resourceCrateDarkRed []byte

	//go:embed assets/CrateDark_Yellow.png
	resourceCrateDarkYellow []byte

	//go:embed assets/EndPoint_Beige.png
	resourceEndpointBeige []byte

	//go:embed assets/EndPoint_Black.png
	resourceEndpointBlack []byte

	//go:embed assets/EndPoint_Blue.png
	resourceEndpointBlue []byte

	//go:embed assets/EndPoint_Brown.png
	resourceEndpointBrown []byte

	//go:embed assets/EndPoint_Gray.png
	resourceEndpointGray []byte

	//go:embed assets/EndPoint_Purple.png
	resourceEndpointPurple []byte

	//go:embed assets/EndPoint_Red.png
	resourceEndpointRed []byte

	//go:embed assets/EndPoint_Yellow.png
	resourceEndpointYellow []byte

	//go:embed assets/Ground_Concrete.png
	resourceGroundConcrete []byte

	//go:embed assets/Ground_Dirt.png
	resourceGroundDirt []byte

	//go:embed assets/Ground_Grass.png
	resourceGroundGrass []byte

	//go:embed assets/Ground_Sand.png
	resourceGroundSand []byte

	//go:embed assets/GroundGravel_Concrete.png
	resourceGroundGravelConcrete []byte

	//go:embed assets/GroundGravel_Dirt.png
	resourceGroundGravelDirt []byte

	//go:embed assets/GroundGravel_Grass.png
	resourceGroundGravelGrass []byte

	//go:embed assets/GroundGravel_Sand.png
	resourceGroundGravelSand []byte

	//go:embed assets/Wall_Beige.png
	resourceWallBeige []byte

	//go:embed assets/Wall_Black.png
	resourceWallBlack []byte

	//go:embed assets/Wall_Brown.png
	resourceWallBrown []byte

	//go:embed assets/Wall_Gray.png
	resourceWallGray []byte

	//go:embed assets/WallRound_Beige.png
	resourceWallRoundBeige []byte

	//go:embed assets/WallRound_Black.png
	resourceWallRoundBlack []byte

	//go:embed assets/WallRound_Brown.png
	resourceWallRoundBrown []byte

	//go:embed assets/WallRound_Gray.png
	resourceWallRoundGray []byte
)

func loadImageFromResource(resource []byte) (*ebiten.Image, error) {
	if nil == resource {
		return nil, fmt.Errorf("Resource nil")
	}

	img, _, err := image.Decode(bytes.NewReader(resource))

	if nil != err {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

func NewSokobanAssets() (*SokobanAssets, error) {
	var err error = nil

	assets := &SokobanAssets{}
	assets.TileWidth = 64
	assets.TileHeight = 64

	// Character
	if assets.ImageCharacter1, err = loadImageFromResource(resourceCharacter1); nil != err {
		return nil, fmt.Errorf("failed loading `Character1': %w", err)
	}

	if assets.ImageCharacter2, err = loadImageFromResource(resourceCharacter2); nil != err {
		return nil, fmt.Errorf("failed loading `Character2': %w", err)
	}

	if assets.ImageCharacter3, err = loadImageFromResource(resourceCharacter3); nil != err {
		return nil, fmt.Errorf("failed loading `Character3': %w", err)
	}

	if assets.ImageCharacter4, err = loadImageFromResource(resourceCharacter4); nil != err {
		return nil, fmt.Errorf("failed loading `Character4': %w", err)
	}

	if assets.ImageCharacter5, err = loadImageFromResource(resourceCharacter5); nil != err {
		return nil, fmt.Errorf("failed loading `Character5': %w", err)
	}

	if assets.ImageCharacter6, err = loadImageFromResource(resourceCharacter6); nil != err {
		return nil, fmt.Errorf("failed loading `Character6': %w", err)
	}

	if assets.ImageCharacter7, err = loadImageFromResource(resourceCharacter7); nil != err {
		return nil, fmt.Errorf("failed loading `Character7': %w", err)
	}

	if assets.ImageCharacter8, err = loadImageFromResource(resourceCharacter8); nil != err {
		return nil, fmt.Errorf("failed loading `Character8': %w", err)
	}

	if assets.ImageCharacter9, err = loadImageFromResource(resourceCharacter9); nil != err {
		return nil, fmt.Errorf("failed loading `Character9': %w", err)
	}

	if assets.ImageCharacter10, err = loadImageFromResource(resourceCharacter10); nil != err {
		return nil, fmt.Errorf("failed loading `Character10': %w", err)
	}

	// Crate
	if assets.ImageCrateBeige, err = loadImageFromResource(resourceCrateBeige); nil != err {
		return nil, fmt.Errorf("failed loading `CrateBeige': %w", err)
	}

	if assets.ImageCrateBlack, err = loadImageFromResource(resourceCrateBlack); nil != err {
		return nil, fmt.Errorf("failed loading `CrateBlack': %w", err)
	}

	if assets.ImageCrateBlue, err = loadImageFromResource(resourceCrateBlue); nil != err {
		return nil, fmt.Errorf("failed loading `CrateBlue': %w", err)
	}

	if assets.ImageCrateBrown, err = loadImageFromResource(resourceCrateBrown); nil != err {
		return nil, fmt.Errorf("failed loading `CrateBrown': %w", err)
	}

	if assets.ImageCrateGray, err = loadImageFromResource(resourceCrateGray); nil != err {
		return nil, fmt.Errorf("failed loading `CrateGray': %w", err)
	}

	if assets.ImageCratePurple, err = loadImageFromResource(resourceCratePurple); nil != err {
		return nil, fmt.Errorf("failed loading `CratePurple': %w", err)
	}

	if assets.ImageCrateRed, err = loadImageFromResource(resourceCrateRed); nil != err {
		return nil, fmt.Errorf("failed loading `CrateRed': %w", err)
	}

	if assets.ImageCrateYellow, err = loadImageFromResource(resourceCrateYellow); nil != err {
		return nil, fmt.Errorf("failed loading `CrateYellow': %w", err)
	}

	// CrateDark
	if assets.ImageCrateDarkBeige, err = loadImageFromResource(resourceCrateDarkBeige); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkBeige': %w", err)
	}

	if assets.ImageCrateDarkBlack, err = loadImageFromResource(resourceCrateDarkBlack); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkBlack': %w", err)
	}

	if assets.ImageCrateDarkBlue, err = loadImageFromResource(resourceCrateDarkBlue); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkBlue': %w", err)
	}

	if assets.ImageCrateDarkBrown, err = loadImageFromResource(resourceCrateDarkBrown); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkBrown': %w", err)
	}

	if assets.ImageCrateDarkGray, err = loadImageFromResource(resourceCrateDarkGray); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkGray': %w", err)
	}

	if assets.ImageCrateDarkPurple, err = loadImageFromResource(resourceCrateDarkPurple); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkPurple': %w", err)
	}

	if assets.ImageCrateDarkRed, err = loadImageFromResource(resourceCrateDarkRed); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkRed': %w", err)
	}

	if assets.ImageCrateDarkYellow, err = loadImageFromResource(resourceCrateDarkYellow); nil != err {
		return nil, fmt.Errorf("failed loading `CrateDarkYellow': %w", err)
	}

	// Endpoint
	if assets.ImageEndpointBeige, err = loadImageFromResource(resourceEndpointBeige); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointBeige': %w", err)
	}

	if assets.ImageEndpointBlack, err = loadImageFromResource(resourceEndpointBlack); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointBlack': %w", err)
	}

	if assets.ImageEndpointBlue, err = loadImageFromResource(resourceEndpointBlue); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointBlue': %w", err)
	}

	if assets.ImageEndpointBrown, err = loadImageFromResource(resourceEndpointBrown); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointBrown': %w", err)
	}

	if assets.ImageEndpointGray, err = loadImageFromResource(resourceEndpointGray); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointGray': %w", err)
	}

	if assets.ImageEndpointPurple, err = loadImageFromResource(resourceEndpointPurple); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointPurple': %w", err)
	}

	if assets.ImageEndpointRed, err = loadImageFromResource(resourceEndpointRed); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointRed': %w", err)
	}

	if assets.ImageEndpointYellow, err = loadImageFromResource(resourceEndpointYellow); nil != err {
		return nil, fmt.Errorf("failed loading `EndpointYellow': %w", err)
	}

	// Ground
	if assets.ImageGroundConcrete, err = loadImageFromResource(resourceGroundConcrete); nil != err {
		return nil, fmt.Errorf("failed loading `GroundConcrete': %w", err)
	}

	if assets.ImageGroundDirt, err = loadImageFromResource(resourceGroundDirt); nil != err {
		return nil, fmt.Errorf("failed loading `GroundDirt': %w", err)
	}

	if assets.ImageGroundGrass, err = loadImageFromResource(resourceGroundGrass); nil != err {
		return nil, fmt.Errorf("failed loading `GroundGrass': %w", err)
	}

	if assets.ImageGroundSand, err = loadImageFromResource(resourceGroundSand); nil != err {
		return nil, fmt.Errorf("failed loading `GroundSand': %w", err)
	}

	// GroundGravel
	if assets.ImageGroundGravelConcrete, err = loadImageFromResource(resourceGroundGravelConcrete); nil != err {
		return nil, fmt.Errorf("failed loading `GroundGravelConcrete': %w", err)
	}

	if assets.ImageGroundGravelDirt, err = loadImageFromResource(resourceGroundGravelDirt); nil != err {
		return nil, fmt.Errorf("failed loading `GroundGravelDirt': %w", err)
	}

	if assets.ImageGroundGravelGrass, err = loadImageFromResource(resourceGroundGravelGrass); nil != err {
		return nil, fmt.Errorf("failed loading `GroundGravelGrass': %w", err)
	}

	if assets.ImageGroundGravelSand, err = loadImageFromResource(resourceGroundGravelSand); nil != err {
		return nil, fmt.Errorf("failed loading `GroundGravelSand': %w", err)
	}

	// Wall
	if assets.ImageWallBeige, err = loadImageFromResource(resourceWallBeige); nil != err {
		return nil, fmt.Errorf("failed loading `WallBeige': %w", err)
	}

	if assets.ImageWallBlack, err = loadImageFromResource(resourceWallBlack); nil != err {
		return nil, fmt.Errorf("failed loading `WallBlack': %w", err)
	}

	if assets.ImageWallBrown, err = loadImageFromResource(resourceWallBrown); nil != err {
		return nil, fmt.Errorf("failed loading `WallBrown': %w", err)
	}

	if assets.ImageWallGray, err = loadImageFromResource(resourceWallGray); nil != err {
		return nil, fmt.Errorf("failed loading `WallGray': %w", err)
	}

	// WallRound
	if assets.ImageWallRoundBeige, err = loadImageFromResource(resourceWallRoundBeige); nil != err {
		return nil, fmt.Errorf("failed loading `WallRoundBeige': %w", err)
	}

	if assets.ImageWallRoundBlack, err = loadImageFromResource(resourceWallRoundBlack); nil != err {
		return nil, fmt.Errorf("failed loading `WallRoundBlack': %w", err)
	}

	if assets.ImageWallRoundBrown, err = loadImageFromResource(resourceWallRoundBrown); nil != err {
		return nil, fmt.Errorf("failed loading `WallRoundBrown': %w", err)
	}

	if assets.ImageWallRoundGray, err = loadImageFromResource(resourceWallRoundGray); nil != err {
		return nil, fmt.Errorf("failed loading `WallRoundGray': %w", err)
	}

	return assets, nil
}
