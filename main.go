package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
)

const (
	canvasSize = 1000
)

type point struct {
	x, y int
}
type Turtle struct {
	position  point
	angle     float64
	isPenDown bool
	penSize   point
	penColor  color.RGBA
	canva     Canva
}

type Canva [canvasSize][canvasSize]color.RGBA
type ElasticCanva [][]color.RGBA

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
		for j := 0; j < t.penSize.x; j++ {
			for k := 0; k < t.penSize.y; k++ {
				t.canva[int(y)+j][int(x)+k] = t.penColor
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

func (t *Turtle) setPenColor(r, g, b uint8) {
	t.penColor.R = r
	t.penColor.G = g
	t.penColor.B = b
}

func (t *Turtle) fill(x int, y int, color color.RGBA) {
	floodFill(x, y, &t.canva, color)
	fmt.Println("Filling at", x, y, "with color", color)

}

// //Flood fill algorithm
func floodFill(x, y int, canvas *Canva, color color.RGBA) {
	if x < 0 || y < 0 || x >= canvasSize || y >= canvasSize {
		return
	}
	if canvas[y][x].A != 0 {
		return
	}
	canvas[y][x] = color
	floodFill(x+1, y, canvas, color)
	floodFill(x-1, y, canvas, color)
	floodFill(x, y+1, canvas, color)
	floodFill(x, y-1, canvas, color)
}

func main() {
	turtle := Turtle{
		position:  point{x: canvasSize / 2, y: canvasSize / 2},
		angle:     0,
		isPenDown: true,
		penSize:   point{x: 1, y: 1},
		penColor:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
		canva:     Canva{},
	}

	// Sample Logo program
	commands := []string{}

	//read from file (first argument)
	filePath := os.Args[1]

	//check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}

	//open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	//read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		commands = append(commands, scanner.Text())
	}

	//execute commands
	for _, command := range commands {
		cmdParts := strings.Split(command, " ")
		cmd := cmdParts[0]

		switch cmd {
		case "fd":
			var distance float64
			fmt.Sscanf(command, "fd %f", &distance)
			turtle.forward(distance)
		case "rt":
			var degrees float64
			fmt.Sscanf(command, "rt %f", &degrees)
			turtle.right(degrees)
		case "lt":
			var degrees float64
			fmt.Sscanf(command, "lt %f", &degrees)
			turtle.left(degrees)
		case "pu":
			turtle.penUp()
		case "pd":
			turtle.penDown()
		case "setpencolor":
			var r, g, b uint8
			fmt.Sscanf(command, "setpencolor [%d %d %d]", &r, &g, &b)
			turtle.setPenColor(r, g, b)
		case "setpensize":
			var sizeX, sizeY int
			fmt.Sscanf(command, "setpensize [%d %d]", &sizeX, &sizeY)
			turtle.setPenSize(sizeX, sizeY)
		case "fill":
			var r, g, b uint8
			fmt.Sscanf(command, "fill [%d %d %d]", &r, &g, &b)
			turtle.fill(turtle.position.x, turtle.position.y, color.RGBA{R: r, G: g, B: b, A: 255})

		}
	}

	// trim canvas according to the minimum and maximum x and y values of the true values
	minX, minY, maxX, maxY := findBoundingBox(turtle.canva)
	// output a png file, with the same name as the input file without the extension (<input>.png)

	//create file
	imageFileName := filePath[:len(filePath)-4] + ".png"
	newImage := ElasticCanva{}
	for y := minY; y <= maxY; y++ {
		row := []color.RGBA{}
		for x := minX; x <= maxX; x++ {
			row = append(row, turtle.canva[y][x])
		}
		newImage = append(newImage, row)
	}
	savePNG(newImage, imageFileName)

	//print a debug output too (content: the array content)
	fileName2 := filePath[:len(filePath)-4] + ".debug.txt"
	file2, err := os.Create(fileName2)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file2.Close()

	for _, row := range newImage {
		for _, pixel := range row {
			if pixel != (color.RGBA{}) {
				file2.WriteString("X")
			} else {
				file2.WriteString(" ")
			}
		}
		file2.WriteString("\n")
	}
	file2.WriteString(fmt.Sprintf("minX: %d, minY: %d, maxX: %d, maxY: %d", minX, minY, maxX, maxY))
	fmt.Println("Debug file saved successfully.")

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

func savePNG(imageArray [][]color.RGBA, filename string) {
	img := image.NewRGBA(image.Rect(0, 0, len(imageArray[0]), len(imageArray)))

	for y, row := range imageArray {
		for x, pixel := range row {
			img.Set(x, y, pixel)
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error encoding PNG:", err)
		return
	}
	fmt.Println("PNG saved successfully.")
}
