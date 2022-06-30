package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	bulletSizeWidth  = 3
	bulletSizeHeight = 10
	bulletSpeed      = 0.06
)

func newBullet(renderer *sdl.Renderer) *element {
	bullet := &element{}

	pr := newPicRenderer(bullet, renderer, "pic/pbullet.bmp")
	bullet.addComponent(pr)

	mover := newBulletMover(bullet, bulletSpeed)
	bullet.addComponent(mover)

	col := circle{
		center: bullet.position,
		radius: 3,
	}
	bullet.collisions = append(bullet.collisions, col)

	bullet.tag = "bullet"

	return bullet
}

var bulletPool []*element

func initBulletPool(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		bul := newBullet(renderer)
		elements = append(elements, bul)
		bulletPool = append(bulletPool, bul)
	}
}

func bulletFromPool() (*element, bool) {
	for _, bul := range bulletPool {
		if !bul.active {
			return bul, true
		}
	}

	return nil, false
}
