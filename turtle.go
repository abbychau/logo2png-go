package main

import (
	"image/color"
	"math"
)

type Turtle struct {
	position  point
	angle     float64
	isPenDown bool
	penSize   point
	penColor  color.RGBA
	canva     Canva
	penType   string // "circle" or "square"
}

func (t *Turtle) forward(distance float64) {
	newX := t.position.x + int(distance*math.Cos(t.angle))
	newY := t.position.y + int(distance*math.Sin(t.angle))

	if t.isPenDown {
		t.drawLine(point{x: newX, y: newY})
	}

	t.position.x = newX
	t.position.y = newY
}

func (t *Turtle) right(degrees float64) {
	t.angle += degrees * (math.Pi / 180)
}

func (t *Turtle) left(degrees float64) {
	t.angle -= degrees * (math.Pi / 180)
}

func (t *Turtle) penUp() {
	t.isPenDown = false
}

func (t *Turtle) penDown() {
	t.isPenDown = true
}

func (t *Turtle) drawLine(end point) {
	dx := end.x - t.position.x
	dy := end.y - t.position.y

	steps := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))

	if steps == 0 {
		return
	}

	xIncrement := int(dx / steps)
	yIncrement := int(dy / steps)

	x := t.position.x
	y := t.position.y

	for i := 0; i < steps; i++ {
		t.canva[int(y)][int(x)] = t.penColor

		// draw a line with the pen size
		if t.penType == "square" {
			for j := 0; j < t.penSize.x; j++ {
				for k := 0; k < t.penSize.y; k++ {
					t.canva[int(y)+j][int(x)+k] = t.penColor
				}
			}
		} else if t.penType == "circle" {
			for j := 0; j < t.penSize.x; j++ {
				for k := 0; k < t.penSize.x; k++ {
					if j*j+k*k < t.penSize.x*t.penSize.x {
						halfJ := j / 2
						halfK := k / 2
						t.canva[y+halfJ][x-halfK] = t.penColor
						t.canva[y-halfJ][x+halfK] = t.penColor
						t.canva[y-halfJ][x-halfK] = t.penColor
						t.canva[y+halfJ][x+halfK] = t.penColor
					}
				}
			}
		}

		x += xIncrement
		y += yIncrement
	}
}
func (t *Turtle) setPenSize(sizeX, sizeY int) {
	t.penSize.x = sizeX
	t.penSize.y = sizeY
}

func (t *Turtle) setPenColor(r, g, b, a uint8) {
	t.penColor.R = r
	t.penColor.G = g
	t.penColor.B = b
	t.penColor.A = a
}

func (t *Turtle) fill(x int, y int, color color.RGBA) {
	if x < 0 || y < 0 || x >= canvasSize || y >= canvasSize {
		return
	}
	if t.canva[y][x].A != 0 {
		return
	}
	t.canva[y][x] = color
	t.fill(x+1, y, color)
	t.fill(x-1, y, color)
	t.fill(x, y+1, color)
	t.fill(x, y-1, color)
}

func findBoundingBox(imageArray Canva) (int, int, int, int) {
	minX := -1
	minY := -1
	maxX := -1
	maxY := -1

	for y, row := range imageArray {
		for x, pixel := range row {
			if pixel != (color.RGBA{}) {
				if minX == -1 {
					minX = x
				}
				if minY == -1 {
					minY = y
				}
				if maxX == -1 {
					maxX = x
				}
				if maxY == -1 {
					maxY = y
				}

				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}
	return minX, minY, maxX, maxY
}
