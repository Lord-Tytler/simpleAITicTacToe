package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func run(renderer *sdl.Renderer, tex *sdl.Texture, pixels []byte) {
	setBackground(white, pixels)
	drawBorders(black, pixels)
	running := true
	playerTurn := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseButtonEvent:
				if t.State == sdl.RELEASED {
					if playerTurn {
						row, col := getSquareClicked(int(t.X), int(t.Y))
						if row != -1 && col != -1 {
							markSquare(row, col, cross)
							drawX(row, col, pixels)
							playerTurn = false
						}
					} else {
						row, col := getSquareClicked(int(t.X), int(t.Y))
						if row != -1 && col != -1 {
							markSquare(row, col, circle)
							drawO(row, col, pixels)
							playerTurn = true
						}
					}
				}
			}
		}
		if checkWin() > 0 {
			fmt.Println("winnenrnernernnrene")
		}
		tex.Update(nil, pixels, winWidth*4)
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)
	}

	return
}

func main() {
	renderer, tex := getGraphics()
	pixels := make([]byte, winWidth*winHeight*4)

	run(renderer, tex, pixels)

	sdl.Delay(2000)

}
