// SPDX-License-Identifier: CC-BY-SA-4.0
package assets

import (
	"log"
	"testing"
)

func TestEmbeddedResources(t *testing.T) {
	if nil == resourceCrateBrown {
		log.Fatalln("Missing embedded resource `resourceCrateBrown'")
	}

	if nil == resourceCrateDarkBrown {
		log.Fatalln("Missing embedded resource `resourceCrateDarkBrown'")
	}

	if nil == resourceCharacter4 {
		log.Fatalln("Missing embedded resource `resourceCharacter4'")
	}

	if nil == resourceEndpointBrown {
		log.Fatalln("Missing embedded resource `resourceEndpointBrown'")
	}

	if nil == resourceGroundGravelConcrete {
		log.Fatalln("Missing embedded resource `resourceGroundGravelConcrete'")
	}

	if nil == resourceWallBrown {
		log.Fatalln("Missing embedded resource `resourceWallBrown'")
	}
}

func TestNewSokobanAssets(t *testing.T) {
	if _, err := NewSokobanAssets(); nil != err {
		log.Fatalf("Failed loading assets: %v\n", err)
	}
}
