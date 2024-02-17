package main

import (
	"image/color"
	"math"
)

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
	floodFill(x, y, &t.canva, color)
}
