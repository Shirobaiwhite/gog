package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const mavicSize = 60

func newMavic(renderer *sdl.Renderer, position vector) *element {
	mavic := &element{}

	mavic.position = position
	mavic.rotation = 180

	pr := newPicRenderer(mavic, renderer, "pic/mavic.bmp")
	mavic.addComponent(pr)

	vtb := newVulnerableToBullets(mavic)
	mavic.addComponent(vtb)

	col := circle{
		center: mavic.position,
		radius: 20,
	}
	mavic.collisions = append(mavic.collisions, col)

	mavic.active = true

	return mavic
}
