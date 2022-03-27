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
	ImageCharacter4           *ebiten.Image
	ImageCrateBrown           *ebiten.Image
	ImageCrateDarkBrown       *ebiten.Image
	ImageEndpointBrown        *ebiten.Image
	ImageGroundGravelConcrete *ebiten.Image
	ImageWallBrown            *ebiten.Image
}

var (
	//go:embed assets/Character4.png
	resourceCharacter4 []byte

	//go:embed assets/Crate_Brown.png
	resourceCrateBrown []byte

	//go:embed assets/CrateDark_Brown.png
	resourceCrateDarkBrown []byte

	//go:embed assets/EndPoint_Brown.png
	resourceEndpointBrown []byte

	//go:embed assets/GroundGravel_Concrete.png
	resourceGroundGravelConcrete []byte

	//go:embed assets/Wall_Brown.png
	resourceWallBrown []byte
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
	assets := &SokobanAssets{}
	var err error = nil

	assets.ImageCrateBrown, err = loadImageFromResource(resourceCrateBrown)

	if nil != err {
		return nil, fmt.Errorf("Failed loading `CrateBrown': %w'", err)
	}

	assets.ImageCrateDarkBrown, err = loadImageFromResource(resourceCrateDarkBrown)

	if nil != err {
		return nil, fmt.Errorf("Failed loading `CrateDarkBrown': %w'", err)
	}

	assets.ImageCharacter4, err = loadImageFromResource(resourceCharacter4)

	if nil != err {
		return nil, fmt.Errorf("Failed loading `Character4': %w'", err)
	}

	assets.ImageEndpointBrown, err = loadImageFromResource(resourceEndpointBrown)

	if nil != err {
		return nil, fmt.Errorf("Failed loading `EndpointBrown': %w'", err)
	}

	assets.ImageGroundGravelConcrete, err = loadImageFromResource(resourceGroundGravelConcrete)

	if nil != err {
		return nil, fmt.Errorf("Failed loading `GroundGravelConcrete': %w'", err)
	}

	assets.ImageWallBrown, err = loadImageFromResource(resourceWallBrown)

	if nil != err {
		return nil, fmt.Errorf("Failed loading `WallBrown': %w'", err)
	}

	return assets, nil
}
