package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600

type color struct { //struct to store color values
	r, g, b byte
}

var black = color{0, 0, 0}
var white = color{255, 255, 255}
var red = color{255, 0, 0}

func setPixel(x, y int, c color, pixels []byte) { // copies the rgb values from color into the respective pixel bytes
	index := (y*winWidth + x) * 4
	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}
func setBackground(c color, pixels []byte) { //sets the background to one solid color
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, c, pixels)
		}
	}
}
func drawBorders(c color, pixels []byte) { //draws lines to divide the canvas into 9 equal parts, leaving whitespace on the outskirts as necessary to ensure perfect squares
	w, h := winWidth, winHeight

	if w < h {
		h = w
	} else if h < w {
		w = h
	}

	wR, hR := (winWidth-w)/2, (winHeight-h)/2
	//fmt.Printf("w: %d   h: %d   wR: %d   hR: %d", w, h, wR, hR)
	y := h/3 + hR
	for x := wR; x < w+wR; x++ {
		setPixel(x, y, c, pixels)
	}
	y = h/3*2 + hR
	for x := wR; x < w+wR; x++ {
		setPixel(x, y, c, pixels)
	}
	x := w/3 + wR
	for y := hR; y < h+hR; y++ {
		setPixel(x, y, c, pixels)
	}
	x = w/3*2 + wR
	for y := hR; y < h+hR; y++ {
		setPixel(x, y, c, pixels)
	}
}

func getSquareClicked(x, y int) (row, col int) {
	halfSize := squareSize() / 2
	row = -1
	col = -1
	if _, cY := getSquareCenter(0, 0); y > cY-halfSize && y < cY+halfSize {
		row = 0
	} else if _, cY := getSquareCenter(1, 0); y > cY-halfSize && y < cY+halfSize {
		row = 1
	} else if _, cY := getSquareCenter(2, 0); y > cY-halfSize && y < cY+halfSize {
		row = 2
	}

	if cX, _ := getSquareCenter(0, 0); x >= cX-halfSize && x <= cX+halfSize {
		col = 0
	} else if cX, _ := getSquareCenter(0, 1); x >= cX-halfSize && x <= cX+halfSize {
		col = 1
	} else if cX, _ := getSquareCenter(0, 2); x >= cX-halfSize && x <= cX+halfSize {
		col = 2
	}
	return row, col
}

func squareSize() int { //calculates the width in pixels of the individual squares
	if winWidth < winHeight {
		return winWidth / 3
	} else {
		return winHeight / 3
	}
}

func getSquareCenter(row, col int) (x int, y int) { //returns the center of a specified square in pixel coordinates
	d := col - 1
	x = winWidth/2 + squareSize()*d
	d = row - 1
	y = winHeight/2 + squareSize()*d
	return x, y
}

func drawX(row, col int, pixels []byte) { //takes center coordinates from getCenterSquare() and draws an x
	size := squareSize()
	sizeModifier := int(.75 * float64(size) * 0.5)
	cX, cY := getSquareCenter(row, col)
	x := cX - sizeModifier
	y := cY - sizeModifier
	for x < cX+sizeModifier && y < cY+sizeModifier {
		setPixel(x, y, red, pixels)
		x++
		y++
	}
	x = cX - sizeModifier
	y = cY + sizeModifier
	for x < cX+sizeModifier && y > cY-sizeModifier {
		setPixel(x, y, red, pixels)
		x++
		y--
	}
}

func drawO(row, col int, pixels []byte) { //takes center coordinates from getCenterSquare and draws a cicle
	size := squareSize()
	r := int(.75 * float64(size) * 0.5)
	cX, cY := getSquareCenter(row, col)
	x := -r
	y := 0
	for x < r {
		y = int(math.Round(math.Sqrt(float64(r*r - x*x))))
		fmt.Println(y)
		setPixel(x+cX, y+cY, red, pixels)
		setPixel(x+cX, (y*-1)+cY, red, pixels)
		x++
	}
}

func getGraphics() (*sdl.Renderer, *sdl.Texture) { //intiialized graphics for the program and returns renderer and texture instances to be used in the main update loop
	window, err := sdl.CreateWindow("TESTING SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
	}

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
	}
	return renderer, tex
}
