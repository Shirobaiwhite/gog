package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type picRenderer struct {
	container *element
	tex       *sdl.Texture

	width, height float64
}

func newPicRenderer(container *element, renderer *sdl.Renderer, filename string) *picRenderer {
	tex := textureFromBMP(renderer, filename)
	_, _, width, height, err := tex.Query()
	if err != nil {
		panic(fmt.Errorf("querying texture: %v", err))
	}
	return &picRenderer{
		container: container,
		tex:       textureFromBMP(renderer, filename),
		width:     float64(width),
		height:    float64(height),
	}
}

func (pr *picRenderer) onDraw(renderer *sdl.Renderer) error {

	x := pr.container.position.x - pr.width/2.0
	y := pr.container.position.y - pr.height/2.0

	renderer.CopyEx(
		pr.tex,
		&sdl.Rect{X: 0, Y: 0, W: int32(pr.width), H: int32(pr.height)},
		&sdl.Rect{X: int32(x), Y: int32(y), W: int32(pr.width), H: int32(pr.height)},
		pr.container.rotation,
		&sdl.Point{X: int32(pr.width) / 2, Y: int32(pr.height) / 2},
		sdl.FLIP_NONE,
	)
	return nil
}

func (pr *picRenderer) onUpdate() error {
	return nil
}

func textureFromBMP(renderer *sdl.Renderer, filename string) *sdl.Texture {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		panic(fmt.Errorf("unable to load pic: %v, %v", filename, err))
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("unable to create texture: %v, %v", filename, err))
	}
	return tex
}

func (pr *picRenderer) onCollision(other *element) error {
	return nil
}
