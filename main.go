package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

const (
	canvasSize = 3000
)

type point struct {
	x, y int
}

type Canva [canvasSize][canvasSize]color.RGBA
type ElasticCanva [][]color.RGBA

func main() {
	turtle := Turtle{
		position:  point{x: canvasSize / 2, y: canvasSize / 2},
		angle:     0,
		isPenDown: true,
		penSize:   point{x: 1, y: 1},
		penColor:  color.RGBA{R: 0, G: 0, B: 0, A: 255},
		canva:     Canva{},
		penType:   "square",
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
		case "bk":
			var distance float64
			fmt.Sscanf(command, "bk %f", &distance)
			turtle.forward(-distance)

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
			var r, g, b, a uint8
			//a is optional
			if len(cmdParts) == 5 {
				fmt.Sscanf(command, "setpencolor [%d %d %d %d]", &r, &g, &b, &a)
			} else {
				fmt.Sscanf(command, "setpencolor [%d %d %d]", &r, &g, &b)
				a = 255
			}
			turtle.setPenColor(r, g, b, a)

		case "setpensize":
			var sizeX, sizeY int
			fmt.Sscanf(command, "setpensize [%d %d]", &sizeX, &sizeY)
			turtle.setPenSize(sizeX, sizeY)
		case "setpentype":
			var penType string
			fmt.Sscanf(command, "setpentype %s", &penType)
			turtle.penType = penType
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
	imageFileName := filePath[:len(filePath)-4] + ".png" //TODO: -4 is a magic number
	newImage := ElasticCanva{}
	for y := minY; y <= maxY; y++ {
		row := []color.RGBA{}
		for x := minX; x <= maxX; x++ {
			row = append(row, turtle.canva[y][x])
		}
		newImage = append(newImage, row)
	}
	savePNG(newImage, imageFileName)

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
	fmt.Println("Image saved to", filename)
}
