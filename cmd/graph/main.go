package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var redColor = color.RGBA{255, 0, 0, 255}      // red
var grayColor = color.RGBA{200, 200, 200, 255} // gray
var pointColor = color.RGBA{0, 0, 0, 255}      // black
var bgColor = color.RGBA{255, 255, 255, 255}   // white

func main() {
	const width, height = 1000, 1000
	const numPoints = 6060
	const angleStep = 5.0
	const distance = 500.0 // Camera distance for perspective

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill background
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, bgColor)
		}
	}

	// generate x y z axis
	for i := 0; i < width; i++ {
		img.Set(i, height/2, pointColor) // x-axis
		img.Set(width/2, i, pointColor)  // y-axis
	}

	for i := 0; i < width; i += width / 10 {
		for j := 0; j < height; j++ {
			if i == width/2 || j == height/2 {
				continue
			}
			img.Set(i, j, grayColor) // x-axis
			img.Set(j, i, grayColor) // y-axis
		}
	}

	vector := []float64{200, 250}
	DrawVector(vector, img)
	// Save to file
	f, err := os.Create("3dplot.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		panic(err)
	}

	println("3D plot saved as 3dplot.png")
}

func GetVectorLength(vector []int) float64 {
	var powed float64 = 0
	for i := range vector {
		powed += math.Pow(float64(vector[i]), 2)
	}

	return math.Sqrt(powed)
}

func DrawVector(vector []float64, img *image.RGBA) {
	factor := float64(vector[1]) / float64(vector[0])
	originX := img.Bounds().Max.X / 2
	originY := img.Bounds().Max.Y / 2
	targetPoint := []int{
		originX + int(vector[0]),
		originY - int(vector[1]),
	}
	fmt.Println("Target Point:", targetPoint)

	for i := 0; i < int(vector[0]); i++ {
		x := originX + i
		y := originY - int(float64(i)*factor)
		img.Set(x, y, redColor)
	}
}
