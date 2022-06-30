package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type keyboardMover struct {
	container *element
	speed     float64

	pr *picRenderer
}

func newkeyboardMover(container *element, speed float64) *keyboardMover {
	return &keyboardMover{
		container: container,
		speed:     speed,
		pr:        container.getComponent(&picRenderer{}).(*picRenderer),
	}
}

func (mover *keyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	cont := mover.container

	if keys[sdl.SCANCODE_LEFT] == 1 {
		// move player to the left
		if cont.position.x >= mover.pr.width/2.0 {
			cont.position.x -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		// move player to the right
		if cont.position.x <= screenWidth-mover.pr.height/2.0 {
			cont.position.x += mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_UP] == 1 {
		// up
		if cont.position.y >= mover.pr.width/2.0 {
			cont.position.y -= mover.speed * delta
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		// down
		if cont.position.y <= screenHeight-mover.pr.height/2.0 {
			cont.position.y += mover.speed * delta
		}
	}

	return nil
}

func (mover *keyboardMover) onCollision(other *element) error {
	return nil
}

func (mover *keyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

type keyboardShooter struct {
	container *element
	cooldown  time.Duration
	lastShot  time.Time
}

func newKeyboardShooter(container *element, cooldown time.Duration) *keyboardShooter {
	return &keyboardShooter{
		container: container,
		cooldown:  cooldown,
	}
}

func (mover *keyboardShooter) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(mover.lastShot) >= mover.cooldown {
			if bul, ok := bulletFromPool(); ok {
				bul.active = true
				bul.position.x = pos.x
				bul.position.y = pos.y - 10
				bul.rotation = 270 * (math.Pi / 180)

				mover.lastShot = time.Now()
			}
		}
	}
	return nil
}

func (mover *keyboardShooter) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *keyboardShooter) onCollision(other *element) error {
	return nil
}
