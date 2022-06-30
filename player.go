package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed  = 0.015
	playerSize   = 60
	playerShotCD = time.Millisecond * 250
)

func newPlayer(renderer *sdl.Renderer) *element {
	player := &element{}
	player.position = vector{
		x: screenWidth / 2.0,
		y: screenHeight - playerSize,
	}

	player.active = true

	pr := newPicRenderer(player, renderer, "pic/player.bmp")
	player.addComponent(pr)

	mover := newkeyboardMover(player, playerSpeed)
	player.addComponent(mover)

	shooter := newKeyboardShooter(player, playerShotCD)
	player.addComponent(shooter)

	return player
}
