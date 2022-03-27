// SPDX-License-Identifier: CC-BY-SA-4.0
package assets

import (
	"log"
	"testing"
)

func TestEmbeddedResources(t *testing.T) {
	// Character
	if nil == resourceCharacter1 {
		log.Fatalln("Missing embedded resource `resourceCharacter1'")
	}

	if nil == resourceCharacter2 {
		log.Fatalln("Missing embedded resource `resourceCharacter2'")
	}

	if nil == resourceCharacter3 {
		log.Fatalln("Missing embedded resource `resourceCharacter3'")
	}

	if nil == resourceCharacter4 {
		log.Fatalln("Missing embedded resource `resourceCharacter4'")
	}

	if nil == resourceCharacter5 {
		log.Fatalln("Missing embedded resource `resourceCharacter5'")
	}

	if nil == resourceCharacter6 {
		log.Fatalln("Missing embedded resource `resourceCharacter6'")
	}

	if nil == resourceCharacter7 {
		log.Fatalln("Missing embedded resource `resourceCharacter7'")
	}

	if nil == resourceCharacter8 {
		log.Fatalln("Missing embedded resource `resourceCharacter8'")
	}

	if nil == resourceCharacter9 {
		log.Fatalln("Missing embedded resource `resourceCharacter9'")
	}

	if nil == resourceCharacter10 {
		log.Fatalln("Missing embedded resource `resourceCharacter10'")
	}

	// Crate
	if nil == resourceCrateBeige {
		log.Fatalln("Missing embedded resource `resourceCrateBeige'")
	}

	if nil == resourceCrateBlack {
		log.Fatalln("Missing embedded resource `resourceCrateBlack'")
	}

	if nil == resourceCrateBlue {
		log.Fatalln("Missing embedded resource `resourceCrateBlue'")
	}

	if nil == resourceCrateBrown {
		log.Fatalln("Missing embedded resource `resourceCrateBrown'")
	}

	if nil == resourceCrateGray {
		log.Fatalln("Missing embedded resource `resourceCrateGray'")
	}

	if nil == resourceCratePurple {
		log.Fatalln("Missing embedded resource `resourceCratePurple'")
	}

	if nil == resourceCrateRed {
		log.Fatalln("Missing embedded resource `resourceCrateRed'")
	}

	if nil == resourceCrateYellow {
		log.Fatalln("Missing embedded resource `resourceCrateYellow'")
	}

	// CrateDark
	if nil == resourceCrateDarkBeige {
		log.Fatalln("Missing embedded resource `resourceCrateDarkBeige'")
	}

	if nil == resourceCrateDarkBlack {
		log.Fatalln("Missing embedded resource `resourceCrateDarkBlack'")
	}

	if nil == resourceCrateDarkBlue {
		log.Fatalln("Missing embedded resource `resourceCrateDarkBlue'")
	}

	if nil == resourceCrateDarkBrown {
		log.Fatalln("Missing embedded resource `resourceCrateDarkBrown'")
	}

	if nil == resourceCrateDarkGray {
		log.Fatalln("Missing embedded resource `resourceCrateDarkGray'")
	}

	if nil == resourceCrateDarkPurple {
		log.Fatalln("Missing embedded resource `resourceCrateDarkPurple'")
	}

	if nil == resourceCrateDarkRed {
		log.Fatalln("Missing embedded resource `resourceCrateDarkRed'")
	}

	if nil == resourceCrateDarkYellow {
		log.Fatalln("Missing embedded resource `resourceCrateDarkYellow'")
	}

	// Endpoint
	if nil == resourceEndpointBeige {
		log.Fatalln("Missing embedded resource `resourceEndpointBeige'")
	}

	if nil == resourceEndpointBlack {
		log.Fatalln("Missing embedded resource `resourceEndpointBlack'")
	}

	if nil == resourceEndpointBlue {
		log.Fatalln("Missing embedded resource `resourceEndpointBlue'")
	}

	if nil == resourceEndpointBrown {
		log.Fatalln("Missing embedded resource `resourceEndpointBrown'")
	}

	if nil == resourceEndpointGray {
		log.Fatalln("Missing embedded resource `resourceEndpointGray'")
	}

	if nil == resourceEndpointPurple {
		log.Fatalln("Missing embedded resource `resourceEndpointPurple'")
	}

	if nil == resourceEndpointRed {
		log.Fatalln("Missing embedded resource `resourceEndpointRed'")
	}

	if nil == resourceEndpointYellow {
		log.Fatalln("Missing embedded resource `resourceEndpointYellow'")
	}

	// Ground
	if nil == resourceGroundConcrete {
		log.Fatalln("Missing embedded resource `resourceGroundConcrete'")
	}

	if nil == resourceGroundDirt {
		log.Fatalln("Missing embedded resource `resourceGroundDirt'")
	}

	if nil == resourceGroundGrass {
		log.Fatalln("Missing embedded resource `resourceGroundGrass'")
	}

	if nil == resourceGroundSand {
		log.Fatalln("Missing embedded resource `resourceGroundSand'")
	}

	// GroundGravel
	if nil == resourceGroundGravelConcrete {
		log.Fatalln("Missing embedded resource `resourceGroundGravelConcrete'")
	}

	if nil == resourceGroundGravelDirt {
		log.Fatalln("Missing embedded resource `resourceGroundGravelDirt'")
	}

	if nil == resourceGroundGravelGrass {
		log.Fatalln("Missing embedded resource `resourceGroundGravelGrass'")
	}

	if nil == resourceGroundGravelSand {
		log.Fatalln("Missing embedded resource `resourceGroundGravelSand'")
	}

	// Wall
	if nil == resourceWallBeige {
		log.Fatalln("Missing embedded resource `resourceWallBeige'")
	}

	if nil == resourceWallBlack {
		log.Fatalln("Missing embedded resource `resourceWallBlack'")
	}

	if nil == resourceWallBrown {
		log.Fatalln("Missing embedded resource `resourceWallBrown'")
	}

	if nil == resourceWallGray {
		log.Fatalln("Missing embedded resource `resourceWallGray'")
	}

	// WallRound
	if nil == resourceWallRoundBeige {
		log.Fatalln("Missing embedded resource `resourceWallRoundBeige'")
	}

	if nil == resourceWallRoundBlack {
		log.Fatalln("Missing embedded resource `resourceWallRoundBlack'")
	}

	if nil == resourceWallRoundBrown {
		log.Fatalln("Missing embedded resource `resourceWallRoundBrown'")
	}

	if nil == resourceWallRoundGray {
		log.Fatalln("Missing embedded resource `resourceWallRoundGray'")
	}
}

func TestNewSokobanAssets(t *testing.T) {
	if _, err := NewSokobanAssets(); nil != err {
		log.Fatalf("Failed loading assets: %v\n", err)
	}
}
