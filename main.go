package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 600
	screenHeight = 600
)

var targetTicksPerSecond = float64(16667)
var second = time.Second
var delta float64

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Shoot'em",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL,
	)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	elements = append(elements, newPlayer(renderer))

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			x := (float64(i)/3)*screenWidth + mavicSize/2.0
			y := float64(j)*mavicSize + mavicSize/2.0

			mavic := newMavic(renderer, vector{x, y})
			elements = append(elements, mavic)
		}
	}

	initBulletPool(renderer)

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, elem := range elements {
			if elem.active {
				err = elem.update()
				if err != nil {
					fmt.Println("updating element:", err)
					return
				}
				err = elem.draw(renderer)
				if err != nil {
					fmt.Println("drawing element:", err)
					return
				}
			}
		}

		if err := checkCollisions(); err != nil {
			fmt.Println("checking collisions:", err)
			return
		}

		renderer.Present()

		targetTicksPerSecond = float64(second / time.Since(frameStartTime))
		// fmt.Println(targetTicksPerSecond)
		delta = float64(time.Since(frameStartTime).Seconds() * targetTicksPerSecond)
	}
}
